package rabbitmq

import (
	"errors"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

//MessageBody is the struct for the body passed in the AMQP message. The type will be set on the Request header
type MessageBody struct {
	Data []byte
	Type string
}

//Message is the amqp request to publish
type Message struct {
	Queue         string
	ReplyTo       string
	ContentType   string
	CorrelationID string
	Priority      uint8
	Body          MessageBody
}

//Connection is the connection created
type Connection struct {
	name     string
	conn     *amqp.Connection
	channel  *amqp.Channel
	exchange string
	queues   []string
	err      chan error
	logras   map[string]*logrus.Logger
}

var (
	connectionPool = make(map[string]*Connection)
	logras         map[string]*logrus.Logger
	log            *logrus.Logger
)

//NewConnection returns the new connection object
func NewConnection(name, exchange string, queues []string, logpool map[string]*logrus.Logger) *Connection {
	if c, ok := connectionPool[name]; ok {
		return c
	}
	c := &Connection{
		exchange: exchange,
		queues:   queues,
		err:      make(chan error),
		logras:   logpool,
	}
	connectionPool[name] = c
	logras = logpool
	log = logpool["info"]
	return c
}

//GetConnection returns the connection which was instantiated
func GetConnection(name string) *Connection {
	return connectionPool[name]
}

// Connect ...
func (c *Connection) Connect() error {
	var err error
	// c.conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	// c.conn, err = amqp.Dial("amqp://localhost:5672")
	// c.conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	// c.conn, err = amqp.Dial("amqp://localhost:5672")
	// c.conn, err = amqp.Dial("amqp://guest:guest@18.140.220.74:5672") 'amqp://guest:guest:18.140.220.74:5672'
	// c.conn, err = amqp.Dial("amqp://guest:guest@13.251.29.17:5672")
	// c.conn, err = amqp.Dial("amqp://guest:guest@18.140.84.228:5672")
	// c.conn, err = amqp.Dial("amqp://guest:guest@52.220.178.251:5672")
	// c.conn, err = amqp.Dial("amqp://ec2-18-140-84-228.ap-southeast-1.compute.amazonaws.com:5672")
	// c.conn, err = amqp.Dial("amqp://guest:guest@18.140.84.228:5672")
	// c.conn, err = amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	// c.conn, err = amqp.Dial("https://imcm.dsldemo.site/rabbitjob/")
	// c.conn, err = amqp.Dial("amqp://guest:guest@rabbit.dsldemo.site/")
	// c.conn, err = amqp.Dial("amqps://rabbit.dsldemo.site") // doesnt work
	// c.conn, err = amqp.Dial("amqp://jain:jain@18.140.84.228:5672") //doesnt works
	// c.conn, err = amqp.Dial("amqp://guest:guest@18.140.84.228:5672") // works well
	// c.conn, err = amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	// c.conn, err = amqp.Dial("amqp://guest:guest@18.140.84.228:5672") // works well
	// c.conn, err = amqp.Dial("amqp://guest:guest@18.140.220.74:5672") //marcus
	// c.conn, err = amqp.Dial("amqp://18.140.220.74:5672") //marcus

	// rabbitURL := "http://rabbit.dsldemo.site"
	// rabbitURL := "rabbit.dsldemo.site"
	// rabbitURL := "amqp://18.140.84.228:5672" //works
	// rabbitURL := "amqp://guest:guest@18.140.84.228:5672" // Jain works
	// rabbitURL := "amqp://guest:guest@localhost:5672/" // Localhost works
	var rabbitURL = ""
	// rabbitURL := "amqp://jain:jain@18.140.220.74:5672//german" //Marcus doesnt work ,error on username passwordnot allowed works after adding vhost
	// RABBIT_URL= "amqp://guest:guest@rabbitmq.aipo-imda.net" works
	rabbitURL = os.Getenv("RABBIT_URL")
	if rabbitURL == "" {
		// rabbitURL = "amqp://jain:jain@18.140.220.74:5672//german"
	}
	log.Info(`RabbitURL:`, rabbitURL)
	c.conn, err = amqp.Dial(rabbitURL)
	if err != nil {
		return err
		// return fmt.Errorf("Error in creating rabbitmq connection with %s : %s", err.Error())
	}
	go func() {
		<-c.conn.NotifyClose(make(chan *amqp.Error)) //Listen to NotifyClose
		c.err <- errors.New("Connection Closed")
	}()
	c.channel, err = c.conn.Channel()
	if err != nil {
		return fmt.Errorf("Channel: %s", err)
	}
	if err := c.channel.ExchangeDeclare(
		c.exchange, //exchange name
		// "my-exchange", //exchange name
		"direct", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // noWait
		nil,      // arguments
	); err != nil {
		log.Error("Error in Exchange Declare ", err)
		return fmt.Errorf("Error in Exchange Declare: %s", err)
	}
	return nil
}

func (c *Connection) BindQueue() error {
	for _, q := range c.queues {
		if _, err := c.channel.QueueDeclare(q, true, false, false, false, nil); err != nil {
			return fmt.Errorf("error in declaring the queue %s", err)
		}
		if err := c.channel.QueueBind(q, "my_routing_key", c.exchange, false, nil); err != nil {
			return fmt.Errorf("Queue  Bind error: %s", err)
		}
	}
	return nil
}

//Reconnect reconnects the connection
func (c *Connection) Reconnect() error {
	if err := c.Connect(); err != nil {
		return err
	}
	if err := c.BindQueue(); err != nil {
		return err
	}
	return nil
}
