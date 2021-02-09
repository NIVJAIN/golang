package models

import (
	"github.com/nivjain/7-ginInterfaceMongoDBRabbitMQ-JWT-RTL/forms"
)

// "go.mongodb.org/mongo-driver/bson"
// "go.mongodb.org/mongo-driver/bson/primitive"

// RabbitRepository ...
type RabbitRepository interface {
	FindByEmailID(EMAILID string) (*forms.UserDetails, error)
	PushJobToRabbitMQ(peopleInfo *forms.PersonDetails) (string, error)
}
