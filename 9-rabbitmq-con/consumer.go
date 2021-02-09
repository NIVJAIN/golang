package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/nivjain/9-rabbitmq/comms"

	"github.com/streadway/amqp"
)

func main() {
	forever := make(chan bool)
	conn := comms.NewConnection("my-consumer-1", "my-exchange", []string{"queue-1", "queue-2"})
	if err := conn.Connect(); err != nil {
		panic(err)
	}
	if err := conn.BindQueue(); err != nil {
		panic(err)
	}
	deliveries, err := conn.Consume()
	if err != nil {
		panic(err)
	}
	for q, d := range deliveries {
		go conn.HandleConsumedDeliveries(q, d, messageHandler)
	}
	<-forever
}

type PersonDetails struct {
	Email string
	Name  string
	Age   int
	City  string
}

func messageHandler(c comms.Connection, q string, deliveries <-chan amqp.Delivery) {
	fmt.Println("messageHandler.....")
	for d := range deliveries {
		m := comms.Message{
			Queue:         q,
			Body:          comms.MessageBody{Data: d.Body, Type: d.Headers["type"].(string)},
			ContentType:   d.ContentType,
			Priority:      d.Priority,
			CorrelationID: d.CorrelationId,
		}
		//handle the custom message
		log.Println("Got message from queue ", m.Queue)
		fmt.Println("gotmessage---", string(m.Body.Data), reflect.TypeOf(m.Body.Data))
		// data, _ := deserialize(m.Body.Data)

		// Parse message.
		// var perDetails PersonDetails
		addTask := &PersonDetails{}
		// json.Unmarshal([]byte(m.Body.Data), &perDetails)
		err := json.Unmarshal(m.Body.Data, addTask)
		if err != nil {
			log.Println(err)
		}
		log.Printf("Result of %s ", addTask.Name)
		fmt.Println("ddd", addTask.Name, addTask.Age, addTask.City, addTask.Email)
		d.Ack(false)
	}
}

type Message map[string]interface{}

func serialize(msg Message) ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(msg)
	return b.Bytes(), err
}

func deserialize(b []byte) (Message, error) {
	var msg Message
	buf := bytes.NewBuffer(b)
	decoder := json.NewDecoder(buf)
	err := decoder.Decode(&msg)
	return msg, err
}
