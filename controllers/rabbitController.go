package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nivjain/7-ginInterfaceMongoDBRabbitMQ-JWT-RTL/forms"
	"github.com/nivjain/7-ginInterfaceMongoDBRabbitMQ-JWT-RTL/models"
)

// //RabbitController ...
// type RabbitController struct {
// 	MClient *mongo.Client
// }

// RabbitHandler will hold everything that controller needs

// RabbitController ...
type RabbitController struct {
	rabbitRepo models.RabbitRepository
}

// SetRabbitController returns a new UserHandler
func (rc *RabbitController) SetRabbitController(rabbitRepo models.RabbitRepository) *RabbitController {
	return &RabbitController{
		rabbitRepo: rabbitRepo,
	}
}

//RabbitFindUser ...
func (rc *RabbitController) RabbitFindUser(c *gin.Context) {
	var user forms.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	getUser, err := rc.rabbitRepo.FindByEmailID(user.Email)
	if err != nil {
		fmt.Println("Error", getUser)
	}
	log.Println("GetUser..", getUser, user)
	c.JSON(http.StatusOK, gin.H{"success": getUser})
}

//PushToRabbitMQ ...
func (rc *RabbitController) PushToRabbitMQ(c *gin.Context) {
	// fmt.Println("iam called...........")
	// var user2 []byte
	var user forms.PersonDetails

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// getUser, err := rc.rabbitRepo.PushJobToRabbitMQ(user.Email)
	getUser, err := rc.rabbitRepo.PushJobToRabbitMQ(&user)

	if err != nil {
		log.Error("Error", getUser)
	}
	// log.Println("GetUser..", getUser, user)
	c.JSON(http.StatusOK, gin.H{"success": getUser})
}
