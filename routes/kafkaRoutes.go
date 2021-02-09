package routes

import (
	"github.com/nivjain/7-ginInterfaceMongoDBRabbitMQ-JWT-RTL/controllers"

	"github.com/gin-gonic/gin"
	"github.com/nivjain/7-ginInterfaceMongoDBRabbitMQ-JWT-RTL/repositories"
	// "github.com/nivjain/7-gininterfaceMongoDB/controllers"
)

// "github.com/nivjain/7-gininterfaceMongoDB/controllers"

// KafkaRouter ...
type KafkaRouter struct{}

var kc = new(controllers.KafkaController)

// KafkaMQHomePage ...
func KafkaMQHomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hi Sripal Jain greetings from rabbitroutes.go homepage",
	})
}

// Routes for user
func (kr KafkaRouter) Routes(mClient *repositories.MongoClient, router *gin.Engine) {
	// rabbitRepo := repositories.SetMongoClient(db, collxpool, rbc, logpool)
	kafkacontroller := kc.SetKafkaController(mClient)
	kafka := router.Group("/kafka")
	kafka.GET("/gethome", KafkaMQHomePage)
	kafka.POST("/getuser", kafkacontroller.KafkaFindUser)
	kafka.POST("/push", kafkacontroller.PushToKafka)
}
