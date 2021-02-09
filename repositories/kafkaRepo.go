package repositories

import (
	"encoding/json"

	"github.com/nivjain/7-ginInterfaceMongoDBRabbitMQ-JWT-RTL/forms"
)

// PushJobToKafka ...
// func (r *MongoClient) PushJobToRabbitMQ(EMAILID string) (*forms.UserDetails, error) {
func (r *MongoClient) PushJobToKafka(peopleInfo *forms.PersonDetails) (string, error) {
	msg, err := json.Marshal(peopleInfo)
	if err != nil {
		return "sd", err
	}
	err = r.kafkaConnection.KafkaPublish(msg)
	if err != nil {
		return "s", err
	}
	return "ss", nil
}
