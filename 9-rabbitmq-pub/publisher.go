package main

import (
	"fmt"

	"github.com/nivjain/9-rabbitmq/comms"
)

func main() {
	conn := comms.NewConnection("my-producer", "my-exchange", []string{"queue-1"})
	if err := conn.Connect(); err != nil {
		panic(err)
	}
	if err := conn.BindQueue(); err != nil {
		panic(err)
	}
	// b := []byte("ABC€")
	var b comms.MessageBody
	b.Data = []byte(`{"hello":"sripaljain","age":"41","City":"Singapore","Id":555}`)
	b.Type = "string"
	m := comms.Message{
		Queue:         "queue-1",
		ReplyTo:       "ssss",
		ContentType:   "string",
		CorrelationID: "id8",
		Priority:      8,
		Body:          b,
		//set the necessary fields
	}
	if err := conn.Publish(m); err != nil {
		panic(err)
	}
	fmt.Println("Successfully delivered the message ....")

}

// for _, q := range c.queues {
// 	m := comms.Message{
// 		Queue:         q,
// 		ReplyTo:       "",
// 		ContentType:   "string",
// 		CorrelationID: "id8",
// 		Priority:      8,
// 		Body:          b,
// 		//set the necessary fields
// 	}
// 	if err := conn.Publish(m); err != nil {
// 		panic(err)
// 	}
// }
