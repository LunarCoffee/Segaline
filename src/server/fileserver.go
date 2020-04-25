package server

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"segaline/src/http"
	"segaline/src/util"
	"strconv"
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
		if req.Method == http.MethodTrace {
			willClose = server.handleTraceRequest(&req, writer)
		} else {
			willClose = server.handleGetOrHeadRequest(&req, writer)
		}

		if willClose {
			break
		}
	}
}

func (server *FileServer) parseRequest(conn net.Conn, writer *bufio.Writer) (req http.Request, ok bool) {
	var err error
	req, err = http.ParseRequest(conn)

	if err == nil {
		switch req.Method {
		case http.MethodGet, http.MethodHead, http.MethodTrace:
			ok = true
		default:
			server.respondErrorTemplate(writer, &req, http.StatusMethodNotAllowed, true)
		}
	} else {
		var status http.StatusCode
		switch err.Error() {
		case util.ErrorContentLengthExceeded:
			status = http.StatusEntityTooLarge
		case util.ErrorRequestURILengthExceeded:
			status = http.StatusRequestURITooLong
		case util.ErrorUnsupportedMethod, util.ErrorUnsupportedTransferEncoding:
			status = http.StatusNotImplemented
		case util.ErrorTimeoutReached:
			status = http.StatusRequestTimeout
		default:
			status = http.StatusBadRequest
		}
		server.respondErrorTemplate(writer, &req, status, true)
	}
	return
}

func (server *FileServer) handleGetOrHeadRequest(req *http.Request, writer *bufio.Writer) bool {
	pathString := req.Uri.PathString()
	if pathString == "/" {
		pathString = util.DefaultEmptyRequestTarget
	}
	filePath := server.fileRoot + pathString
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		server.respondErrorTemplate(writer, req, http.StatusNotFound, false)
		return false
	}
	contentType := server.contentTypeByExt(pathString[strings.LastIndex(pathString, ".")+1:])

	res := http.NewResponse(req).WithStatus(http.StatusOK)
	if len(content) <= util.ResponseMaxUnchunkedBody {
		res.WithHeader(http.HeaderETag, "\"" + getETag(content) + "\"")
	}
	if info, err := os.Stat(filePath); err == nil {
		res.WithHeader(http.HeaderLastModified, formatTimeGMT(info.ModTime()))
	}
	if req.Method == http.MethodGet {
		res.WithBody(content, contentType)
	}

	res.Respond(writer)
	return req.WillCloseConnection()
}

func (*FileServer) handleTraceRequest(req *http.Request, writer *bufio.Writer) bool {
	http.NewResponse(req).WithStatus(http.StatusOK).WithBody(req.AsBytes(), http.MediaTypeHTTP).Respond(writer)
	return req.WillCloseConnection()
}

func (server *FileServer) respondErrorTemplate(
	writer *bufio.Writer,
	req *http.Request,
	status http.StatusCode,
	close bool,
) {
	template, err := ioutil.ReadFile(server.templateRoot + "/error.html")
	content := template
	if err != nil {
		content = []byte(util.DefaultFallbackErrorTemplate)
	}
	content = []byte(server.formatErrorTemplate(string(content), status))

	res := http.NewResponse(req).WithStatus(status).WithBody(content, http.MediaTypeHTML)
	if close {
		res.WithHeader(http.HeaderConnection, string(http.ConnectionHeaderClose))
	}
	if status == http.StatusMethodNotAllowed {
		res.WithHeader(http.HeaderAllow, fmt.Sprintf("%s, %s, %s", http.MethodGet, http.MethodHead, http.MethodTrace))
	}
	res.Respond(writer)
}

func (*FileServer) formatErrorTemplate(template string, status http.StatusCode) string {
	statusReplaced := strings.ReplaceAll(template, "{statusCode}", strconv.Itoa(int(status)))
	return strings.ReplaceAll(statusReplaced, "{serverInfo}", util.ServerNameVersion)
}

func (*FileServer) closeConnectionLog(conn net.Conn) {
	if err := conn.Close(); err != nil {
		log.Println("An issue occurred while closing a client connection.")
	}
}

func (*FileServer) contentTypeByExt(ext string) http.MediaType {
	switch ext {
	case "aac":
		return http.MediaTypeAAC
	case "avi":
		return http.MediaTypeAVI
	case "bmp":
		return http.MediaTypeBitmap
	case "css":
		return http.MediaTypeCSS
	case "csv":
		return http.MediaTypeCSV
	case "epub":
		return http.MediaTypeEPUB
	case "gz":
		return http.MediaTypeGZip
	case "gif":
		return http.MediaTypeGIF
	case "htm", "html":
		return http.MediaTypeHTML
	case "ico":
		return http.MediaTypeIcon
	case "jpg", "jpeg":
		return http.MediaTypeJPEG
	case "js":
		return http.MediaTypeJavaScript
	case "json":
		return http.MediaTypeJSON
	case "mp3":
		return http.MediaTypeMP3
	case "mp4":
		return http.MediaTypeMP4
	case "oga":
		return http.MediaTypeOGGAudio
	case "png":
		return http.MediaTypePNG
	case "pdf":
		return http.MediaTypePDF
	case "php":
		return http.MediaTypePHP
	case "rtf":
		return http.MediaTypeRTF
	case "svg":
		return http.MediaTypeSVG
	case "swf":
		return http.MediaTypeSWF
	case "ttf":
		return http.MediaTypeTTF
	case "txt":
		return http.MediaTypeText
	case "wav":
		return http.MediaTypeWAV
	case "weba":
		return http.MediaTypeWEBMAudio
	case "webm":
		return http.MediaTypeWEBMVideo
	case "webp":
		return http.MediaTypeWEBPImage
	case "woff":
		return http.MediaTypeWOFF
	case "woff2":
		return http.MediaTypeWOFF2
	case "xhtml":
		return http.MediaTypeXHTML
	case "xml":
		return http.MediaTypeXML
	case "zip":
		return http.MediaTypeZip
	}
	return http.MediaTypeBinary
}
