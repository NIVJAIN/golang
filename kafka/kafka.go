package kafka

import (
	"fmt"
	"os"

	// "log"
	logams "log"
	// "github.com/nivjain/7-ginInterfaceMongoDBRabbitMQ-JWT-RTL/middlewares/logger"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

const (
	// kafkaConn = "10.4.1.29:9092"
	kafkaConn = "127.0.0.1:9092"
	topic     = "senz"
)

// Connection ...
type Connection struct {
	KafkaProducer sarama.SyncProducer
	topic         string
}

// InitProducer kafka publisher
func (c *Connection) KafkaInitConnection(setTopic string, logsetup *logrus.Logger) (*Connection, error) {
	// setup sarama log to stdout
	// sarama.Logger = log.New(os.Stdout, "", log.Ltime)
	sarama.Logger = logams.New(os.Stdout, "", logams.Ltime)
	log = logsetup
	// producer config
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	// async producer
	//prd, err := sarama.NewAsyncProducer([]string{kafkaConn}, config)

	// sync producer
	prd, err := sarama.NewSyncProducer([]string{kafkaConn}, config)
	if err != nil {
		return nil, err
	}
	var con Connection
	con.KafkaProducer = prd
	con.topic = setTopic
	log.Info("Kafka Connection Succesfull...")
	return &con, nil
	// return prd, err
}

// KafkaPublish kafka messages to the broker
func (c *Connection) KafkaPublish(message []byte) error {
	// publish sync
	msg := &sarama.ProducerMessage{
		Topic: c.topic,
		// Value: sarama.StringEncoder(message),
		Value: sarama.ByteEncoder(message),
	}
	p, o, err := c.KafkaProducer.SendMessage(msg)
	if err != nil {
		log.Println("Error publish: ", err.Error())
		return err
	}

	// publish async
	//producer.Input() <- &sarama.ProducerMessage{

	log.Println("Partition: ", p)
	log.Println("Offset: ", o)
	return nil
}

func (c *Connection) publish(message string, producer sarama.SyncProducer) {
	// publish sync
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	p, o, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("Error publish: ", err.Error())
	}

	// publish async
	//producer.Input() <- &sarama.ProducerMessage{

	fmt.Println("Partition: ", p)
	fmt.Println("Offset: ", o)
}
