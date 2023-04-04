package transport

func (s *Server) routes() {
	s.HTTP.POST("/users", s.h.User.Create)
}
