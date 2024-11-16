package api

func Setup(s *Server) {
	tickets := s.Router.Group("/tickets")
	{
		tickets.GET("/ticket", s.GetTickets)
		tickets.POST("/ticket", s.CreateTicket)
		tickets.PATCH("/ticket/:id", s.UpdateTicket)
		tickets.DELETE("/ticket/:id", s.DeleteTicket)
	}

	orders := s.Router.Group("/orders")
	{
		orders.POST("/order", s.CreateOrder)
		orders.GET("/order/:id", s.GetOrder)
		// TODO: orders.POST("/order/:id ", s.PayOrder)
		// TODO: orders.POST("/order/:id", s.DeleteOrder)
	}
	// TODO: client := s.Router.Group("/clients"); {}
}
