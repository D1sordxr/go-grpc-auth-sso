package api

func Setup(s *Server) {
	api := s.Router.Group("/tickets")
	{
		api.GET("/ticket", s.GetTicket)
		api.POST("/ticket", s.CreateTicket)
	}
}
