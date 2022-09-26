package v1

import (
	"SSEHandler/ssehandler"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	sseHandler := ssehandler.NewSSEHandler()
	sseHandler.HandleEvents()

	h := handler.Group("/api/v1")
	{
		handlerController := New(sseHandler)
		handlerController.newUserRoutes(h)
	}
}
