package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nivjain/7-ginInterfaceMongoDBRabbitMQ-JWT-RTL/forms"
	"github.com/nivjain/7-ginInterfaceMongoDBRabbitMQ-JWT-RTL/models"
)

// KafkaHandler will hold everything that controller needs

// KafkaController ...
type KafkaController struct {
	kafkaRepo models.KafkaRepository
}

// SetKafkaController returns a new UserHandler
func (rc *KafkaController) SetKafkaController(kafkaRepo models.KafkaRepository) *KafkaController {
	return &KafkaController{
		kafkaRepo: kafkaRepo,
	}
}

// KafkaFindUser ...
func (rc *KafkaController) KafkaFindUser(c *gin.Context) {
	// var user forms.User
	// if err := c.ShouldBindJSON(&user); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	// getUser, err := rc.rabbitRepo.FindByEmailID(user.Email)
	// if err != nil {
	// 	fmt.Println("Error", getUser)
	// }
	// log.Println("GetUser..", getUser, user)
	// c.JSON(http.StatusOK, gin.H{"success": getUser})
}

// PushToKafka ...
func (rc *KafkaController) PushToKafka(c *gin.Context) {
	// fmt.Println("iam called...........")
	// var user2 []byte
	var user forms.PersonDetails

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// getUser, err := rc.rabbitRepo.PushJobToRabbitMQ(user.Email)
	getUser, err := rc.kafkaRepo.PushJobToKafka(&user)

	if err != nil {
		log.Error("Error", getUser)
	}
	// log.Println("GetUser..", getUser, user)
	c.JSON(http.StatusOK, gin.H{"success": getUser})
}
