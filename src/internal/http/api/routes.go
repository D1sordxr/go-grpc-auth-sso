package api

func Setup(s *Server) {
	api := s.Router.Group("/tickets")
	{
		api.GET("/ticket", s.GetTicket)
		api.POST("/ticket", s.CreateTicket)
		api.PATCH("/ticket/:id", s.UpdateTicket)
		api.DELETE("/ticket/:id", s.DeleteTicket)
	}
}
