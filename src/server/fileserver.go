package server

import (
	"bufio"
	"io/ioutil"
	"log"
	"net"
	"segaline/src/request"
	"segaline/src/response"
	"segaline/src/util"
	"strings"
)

type FileServer struct {
	listener   net.Listener
	acceptChan chan net.Conn

	fileRoot string
}

func NewFileServer(fileRoot string) Server {
	return &FileServer{
		acceptChan: make(chan net.Conn),
		fileRoot:   strings.TrimSuffix(fileRoot, "/"),
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
	defer closeConnectionLog(conn)
	writer := bufio.NewWriterSize(conn, util.ResponseWriterBufferSize)

	for req, ok := parseRequest(conn, writer); ok; req, ok = parseRequest(conn, writer) {
		pathString := req.Uri.PathString()
		if pathString == "/" {
			pathString = util.DefaultEmptyRequestTarget
		}
		content, err := ioutil.ReadFile(server.fileRoot + pathString)
		if err != nil {
			response.RespondStatus(writer, util.HttpStatusNotFound, false)
			continue
		}

		contentType := util.ContentTypeByExt(pathString[strings.LastIndex(pathString, ".")+1:])
		res := response.New().WithStatus(util.HttpStatusOK)
		if req.Method == util.HttpMethodGet {
			res.WithBody(content, contentType)
		}

		res.Respond(writer)
		log.Printf("%s %s %s\n", req.Method, &req.Uri, conn.RemoteAddr())

		if req.WillCloseConnection() {
			break
		}
	}
}

func parseRequest(conn net.Conn, writer *bufio.Writer) (req request.Request, ok bool) {
	var err error
	req, err = request.Parse(conn)

	if err == nil {
		if req.Method != util.HttpMethodGet && req.Method != util.HttpMethodHead {
			response.RespondStatus(writer, util.HttpStatusMethodNotAllowed, true)
		} else {
			ok = true
		}
	} else {
		var status util.HttpStatusCode
		switch err.Error() {
		case util.ErrorContentLengthExceeded:
			status = util.HttpStatusPayloadTooLarge
		case util.ErrorRequestURILengthExceeded:
			status = util.HttpStatusRequestURITooLong
		case util.ErrorUnsupportedTransferEncoding:
			status = util.HttpStatusNotImplemented
		case util.ErrorTimeoutReached:
			status = util.HttpStatusRequestTimeout
		default:
			status = util.HttpStatusBadRequest
		}
		response.RespondStatus(writer, status, true)
	}
	return
}

func closeConnectionLog(conn net.Conn) {
	if err := conn.Close(); err != nil {
		log.Println("An issue occurred while closing a client connection.")
	}
}
