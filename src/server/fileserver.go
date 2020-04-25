package server

import (
	"bufio"
	"io/ioutil"
	"log"
	"net"
	"os"
	"segaline/src/request"
	"segaline/src/response"
	"segaline/src/util"
	"strings"
)

type FileServer struct {
	listener   net.Listener
	acceptChan chan net.Conn

	fileRoot     string
	templateRoot string
}

func NewFileServer(fileRoot string, templateRoot string) Server {
	return &FileServer{
		acceptChan:   make(chan net.Conn),
		fileRoot:     strings.TrimSuffix(fileRoot, "/"),
		templateRoot: strings.TrimSuffix(templateRoot, "/"),
	}
}

func (server *FileServer) Start(addr string) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	server.listener = listener
	go func() {
		for {
			conn, err := server.listener.Accept()
			if err != nil {
				close(server.acceptChan)
				break
			}
			server.acceptChan <- conn
		}
	}()

	for conn := range server.acceptChan {
		go server.handleClient(conn)
	}
	return nil
}

func (server *FileServer) Stop() error {
	return server.listener.Close()
}

func (server *FileServer) handleClient(conn net.Conn) {
	defer server.closeConnectionLog(conn)
	writer := bufio.NewWriterSize(conn, util.ResponseWriterBufferSize)

	for req, ok := server.parseRequest(conn, writer); ok; req, ok = server.parseRequest(conn, writer) {
		var willClose bool
		if req.Method == util.MethodTrace {
			willClose = server.handleTrace(&req, writer)
		} else {
			willClose = server.handleGetOrHead(&req, writer)
		}

		log.Printf("%s %s %s\n", req.Method, &req.Uri, conn.RemoteAddr())
		if willClose {
			break
		}
	}
}

func (server *FileServer) handleGetOrHead(req *request.Request, writer *bufio.Writer) bool {
	pathString := req.Uri.PathString()
	if pathString == "/" {
		pathString = util.DefaultEmptyRequestTarget
	}
	filePath := server.fileRoot + pathString
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		server.respondErrorTemplate(writer, util.StatusNotFound, false)
		return false
	}
	contentType := util.ContentTypeByExt(pathString[strings.LastIndex(pathString, ".")+1:])

	res := response.New().WithStatus(util.StatusOK)
	if len(content) <= util.ResponseMaxUnchunkedBody {
		res.WithHeader(util.HeaderETag, util.GetETag(content))
	}
	if info, err := os.Stat(filePath); err == nil {
		res.WithHeader(util.HeaderLastModified, util.FormatTimeGMT(info.ModTime()))
	}
	if req.Method == util.MethodGet {
		res.WithBody(content, contentType)
	}

	res.Respond(writer)
	return req.WillCloseConnection()
}

func (*FileServer) handleTrace(req *request.Request, writer *bufio.Writer) bool {
	response.New().WithStatus(util.StatusOK).WithBody(req.AsBytes(), util.MediaTypeHTTP).Respond(writer)
	return req.WillCloseConnection()
}

func (server *FileServer) parseRequest(conn net.Conn, writer *bufio.Writer) (req request.Request, ok bool) {
	var err error
	req, err = request.Parse(conn)

	if err == nil {
		switch req.Method {
		case util.MethodGet, util.MethodHead, util.MethodTrace:
			ok = true
		default:
			server.respondErrorTemplate(writer, util.StatusMethodNotAllowed, true)
		}
	} else {
		var status util.HttpStatusCode
		switch err.Error() {
		case util.ErrorContentLengthExceeded:
			status = util.StatusEntityTooLarge
		case util.ErrorRequestURILengthExceeded:
			status = util.StatusRequestURITooLong
		case util.ErrorUnsupportedTransferEncoding:
			status = util.StatusNotImplemented
		case util.ErrorTimeoutReached:
			status = util.StatusRequestTimeout
		default:
			status = util.StatusBadRequest
		}
		server.respondErrorTemplate(writer, status, true)
	}
	return
}

func (server *FileServer) respondErrorTemplate(writer *bufio.Writer, status util.HttpStatusCode, close bool) {
	template, err := ioutil.ReadFile(server.templateRoot + "/error.html")
	content := template
	if err != nil {
		content = []byte(util.DefaultFallbackErrorTemplate)
	}
	content = []byte(util.FormatErrorTemplate(string(content), status))

	res := response.New().WithStatus(status).WithBody(content, util.MediaTypeHTML)
	if close {
		res.WithHeader(util.HeaderConnection, string(util.ConnectionClose))
	}
	res.Respond(writer)
}

func (*FileServer) closeConnectionLog(conn net.Conn) {
	if err := conn.Close(); err != nil {
		log.Println("An issue occurred while closing a client connection.")
	}
}
