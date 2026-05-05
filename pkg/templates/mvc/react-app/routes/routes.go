package routes

import (
	"github.com/gin-gonic/gin"
	"react-app/internal/handlers"
)

func Setup() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/users", handlers.GetUsers)
	}
	return r
}
