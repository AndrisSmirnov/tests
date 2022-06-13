package http_controllers

import (
	"chat/api/dto"
	middlewares "chat/app/controllers/middleware"

	"net/http"

	"github.com/gin-gonic/gin"
)

func ChatController(router *gin.Engine) {
	authGroup := router.Group("/chat")
	{
		authGroup.POST("/sendMessage", middlewares.Validator(dto.SendMessageRequest{}), func(c *gin.Context) {
			data := c.MustGet("validData").(*dto.SendMessageRequest)

			c.JSON(http.StatusOK, data)
		})
	}
}
