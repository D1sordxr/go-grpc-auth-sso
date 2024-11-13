package api

import "github.com/gin-gonic/gin"

func Setup(r *gin.Engine) {
	api := r.Group("/tickets")
	{
		api.GET("/ticket", GetTicket)
		api.POST("/ticket", CreateTicket)
	}
}
