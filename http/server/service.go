package server

// Service defines a server service which can self register
type Service interface {
	Register(server *Server)
}
