package server

// Option defines a function which applies configuration to server
type Option func(*Server)

// WithAddr appends the address of the server
func WithAddr(a string) Option {
	return func(s *Server) {
		s.srv.Addr = a
	}
}