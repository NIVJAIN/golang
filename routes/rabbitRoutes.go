package routes

import (
	"github.com/nivjain/7-ginInterfaceMongoDBRabbitMQ-JWT-RTL/controllers"
	"github.com/nivjain/7-ginInterfaceMongoDBRabbitMQ-JWT-RTL/rabbitmq"
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/nivjain/7-ginInterfaceMongoDBRabbitMQ-JWT-RTL/repositories"
	"go.mongodb.org/mongo-driver/mongo"
	// "github.com/nivjain/7-gininterfaceMongoDB/controllers"
)

// "github.com/nivjain/7-gininterfaceMongoDB/controllers"

//RabbitRouter ...
type RabbitRouter struct{}

var rc = new(controllers.RabbitController)

// var Usercontroller = new(controllers.UserController)

// RabbitMQHomePage ...
func RabbitMQHomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hi Sripal Jain greetings from rabbitroutes.go homepage",
	})
}

// Routes for user
func (rbr RabbitRouter) Routes(db *mongo.Client, collxpool map[string]*mongo.Collection, rbc *rabbitmq.Connection, logpool map[string]*logrus.Logger, mClient *repositories.MongoClient, router *gin.Engine) {
	// rabbitRepo := repositories.SetMongoClient(db, collxpool, rbc, logpool)
	rbcontroller := rc.SetRabbitController(mClient)
	rabbit := router.Group("/rabbit")
	rabbit.GET("/gethome", RabbitMQHomePage)
	rabbit.POST("/getuser", rbcontroller.RabbitFindUser)
	rabbit.POST("/push", rbcontroller.PushToRabbitMQ)
}
