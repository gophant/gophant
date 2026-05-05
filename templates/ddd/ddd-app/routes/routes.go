package routes

import (
	dhttp "ddd-app/interfaces/http"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/users", dhttp.GetUsers)
	}
	return r
}
