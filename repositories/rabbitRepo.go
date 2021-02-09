package repositories

import (
	"encoding/json"

	"github.com/nivjain/7-ginInterfaceMongoDBRabbitMQ-JWT-RTL/forms"
	"github.com/nivjain/7-ginInterfaceMongoDBRabbitMQ-JWT-RTL/rabbitmq"
)

// PushJobToRabbitMQ ...
// func (r *MongoClient) PushJobToRabbitMQ(EMAILID string) (*forms.UserDetails, error) {
func (r *MongoClient) PushJobToRabbitMQ(peopleInfo *forms.PersonDetails) (string, error) {
	var b rabbitmq.MessageBody
	// b.Data = []byte(`{"hello":"sripaljain","age":"41","City":"Singapore","Id":555}`)
	// s, _ := json.Marshal(EMAILID)
	// b.Data = s
	// addTask := forms.PersonDetails{Name: "sripaljain", Age: 41, City: "singapore", Email: "sripal.jain@gmail.com"}
	s, _ := json.Marshal(peopleInfo)
	b.Data = s
	b.Type = "string"
	m := rabbitmq.Message{
		Queue:   "queue-1",
		ReplyTo: "",
		// ContentType:   "string",
		ContentType:   "text/plain",
		CorrelationID: "id8",
		Priority:      8,
		Body:          b,
		//set the necessary fields
	}
	if err := r.rabbitmqConnection.Publish(m); err != nil {
		panic(err)
	}
	return "RabbitMQ::::message published succesfully", nil
}
