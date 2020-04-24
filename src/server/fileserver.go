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
	writer := bufio.NewWriter(conn)
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("An unexpected issue occurred while closing a client connection.")
		}
	}()

	req, err := request.Parse(conn)
	if err != nil {
		respondStatus(writer, util.HttpStatusBadRequest)
		return
	}

	pathString := req.Uri.PathString()
	content, err := ioutil.ReadFile(server.fileRoot + pathString)
	if err != nil {
		respondStatus(writer, util.HttpStatusNotFound)
		return
	}

	log.Printf("%s %s %s\n", req.Method, &req.Uri, conn.RemoteAddr())

	contentType := contentTypeByExt(pathString[strings.LastIndex(pathString, ".")+1:])
	res := response.New().WithBody(content, contentType)
	respond(writer, res)
}

func contentTypeByExt(ext string) util.HttpMediaType {
	switch ext {
	case "aac":
		return util.HttpMediaTypeAAC
	case "avi":
		return util.HttpMediaTypeAVI
	case "bmp":
		return util.HttpMediaTypeBitmap
	case "css":
		return util.HttpMediaTypeCSS
	case "csv":
		return util.HttpMediaTypeCSV
	case "epub":
		return util.HttpMediaTypeEPUB
	case "gz":
		return util.HttpMediaTypeGZip
	case "gif":
		return util.HttpMediaTypeGIF
	case "htm", "html":
		return util.HttpMediaTypeHTML
	case "ico":
		return util.HttpMediaTypeIcon
	case "jpg", "jpeg":
		return util.HttpMediaTypeJPEG
	case "js":
		return util.HttpMediaTypeJavaScript
	case "json":
		return util.HttpMediaTypeJSON
	case "mp3":
		return util.HttpMediaTypeMP3
	case "mp4":
		return util.HttpMediaTypeMP4
	case "oga":
		return util.HttpMediaTypeOGGAudio
	case "png":
		return util.HttpMediaTypePNG
	case "pdf":
		return util.HttpMediaTypePDF
	case "php":
		return util.HttpMediaTypePHP
	case "rtf":
		return util.HttpMediaTypeRTF
	case "svg":
		return util.HttpMediaTypeSVG
	case "swf":
		return util.HttpMediaTypeSWF
	case "ttf":
		return util.HttpMediaTypeTTF
	case "txt":
		return util.HttpMediaTypeText
	case "wav":
		return util.HttpMediaTypeWAV
	case "weba":
		return util.HttpMediaTypeWEBMAudio
	case "webm":
		return util.HttpMediaTypeWEBMVideo
	case "webp":
		return util.HttpMediaTypeWEBPImage
	case "woff":
		return util.HttpMediaTypeWOFF
	case "woff2":
		return util.HttpMediaTypeWOFF2
	case "xhtml":
		return util.HttpMediaTypeXHTML
	case "xml":
		return util.HttpMediaTypeXML
	case "zip":
		return util.HttpMediaTypeZip
	}
	return util.HttpMediaTypeBinary
}

func respondStatus(writer *bufio.Writer, status util.HttpStatusCode) {
	respond(writer, response.New().WithStatus(status))
}

func respond(writer *bufio.Writer, res *response.Response) {
	if err := writeFully(writer, res.AsBytes()); err != nil {
		log.Println("An issue occurred while responding to a request.")
	}
}

func writeFully(writer *bufio.Writer, bytes []byte) error {
	written := 0
	for written < len(bytes) {
		n, err := writer.Write(bytes[written:])
		if err != nil {
			return err
		}
		written += n
	}
	return writer.Flush()
}
