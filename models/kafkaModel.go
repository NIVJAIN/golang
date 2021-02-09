package models

import (
	"github.com/nivjain/7-ginInterfaceMongoDBRabbitMQ-JWT-RTL/forms"
)

// "go.mongodb.org/mongo-driver/bson"
// "go.mongodb.org/mongo-driver/bson/primitive"

// KafkaRepository ...
type KafkaRepository interface {
	FindByEmailID(EMAILID string) (*forms.UserDetails, error)
	PushJobToKafka(peopleInfo *forms.PersonDetails) (string, error)
}
