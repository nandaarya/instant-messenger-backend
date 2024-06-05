package routes

import (
	"instant-messenger-backend/controllers"
	"instant-messenger-backend/rabbitmq"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
  r.POST("/messages", rabbitmq.PublishToDatabase)
  r.POST("/chatList", controllers.GetChatList)
  r.POST("/chatDetail", controllers.GetChatDetail)

  r.GET("/", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "Hello, world! Connect to Instant Messenger App!",
    })
  })

}