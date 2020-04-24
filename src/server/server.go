package server

type Server interface {
	Start(addr string) error
	Stop() error
}
