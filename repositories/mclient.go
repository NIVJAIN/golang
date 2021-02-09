package repositories

import (
	"github.com/nivjain/7-ginInterfaceMongoDBRabbitMQ-JWT-RTL/kafka"
	"github.com/nivjain/7-ginInterfaceMongoDBRabbitMQ-JWT-RTL/rabbitmq"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoClient implements models.UserRepository
type MongoClient struct {
	db                 *mongo.Client
	collection         map[string]*mongo.Collection
	rabbitmqConnection *rabbitmq.Connection
	logcollections     map[string]*logrus.Logger
	kafkaConnection    *kafka.Connection
}

func SetLogCollection(logpool map[string]*logrus.Logger) *MongoClient {
	// lp := make(map[string]*logrus.Logger)
	return &MongoClient{
		logcollections: logpool,
	}
}

// SetMongoClient ..
func SetMongoClient(db *mongo.Client, collx map[string]*mongo.Collection, rbc *rabbitmq.Connection, logpool map[string]*logrus.Logger, kafkaConn *kafka.Connection) *MongoClient {
	log = logpool["info"]
	return &MongoClient{
		db:                 db,
		collection:         collx,
		rabbitmqConnection: rbc,
		logcollections:     logpool,
		kafkaConnection:    kafkaConn,
	}
}

// GetMongoClient ...
func (m *MongoClient) GetMongoClient() *MongoClient {
	return m
}
