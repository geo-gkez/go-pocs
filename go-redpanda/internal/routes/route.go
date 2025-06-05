package routes

import (
	"net/http"

	"github.com/geo-gkez/go-pocs/redpanda-poc/internal/model"
	"github.com/geo-gkez/go-pocs/redpanda-poc/internal/service"
	"github.com/gin-gonic/gin"
)

func SetupRoutesAndRegister(router *gin.Engine, kafka service.IKafkaService) *gin.Engine {
	// Define the routes for the application
	router.POST("/produce", func(c *gin.Context) {
		produceMessage(c, kafka)
	})
	return router
}

func produceMessage(ctx *gin.Context, kafkaService service.IKafkaService) {
	var message model.ProduceMessageRequest

	if err := ctx.ShouldBindJSON(&message); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	kafkaService.ProduceMessage(message)

	ctx.JSON(202, gin.H{
		"message": "Message produced successfully!",
	})

}
