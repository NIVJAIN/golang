package db

import (
	"context"
	"fmt"

	// "log"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	"github.com/nivjain/7-ginInterfaceMongoDBRabbitMQ-JWT-RTL/middlewares/logger"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// 	"strconv"
)

//MongoStruct ...
type MongoStruct struct {
	MongoClient *mongo.Client
	Ctx         context.Context
	Cancel      context.CancelFunc
}

//MONGOClient ...
var MONGOClient *mongo.Client

// var logras *logrus.Logger
var log *logrus.Logger

const (
	// Timeout operations after N seconds
	connectTimeout           = 5
	connectionStringTemplate = "mongodb://%s:%s@%s"
)

func init() {
	// logger.Log.Println("Init called from From MOngoDB....")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file, please create one in the root directory")
	}
	var logams = new(logger.Logams)
	log = logams.SetLogger("info")
}

//ConnectDB ...
func ConnectDB(l map[string]*logrus.Logger) (*mongo.Client, context.Context, context.CancelFunc) {
	username := os.Getenv("MONGODB_USERNAME")
	password := os.Getenv("MONGODB_PASSWORD")
	clusterEndpoint := os.Getenv("MONGODB_ENDPOINT")
	// MONGODB_USERNAME="root"
	// MONGODB_PASSWORD="$iloveblockchain"
	// MONGODB_ENDPOINT="localhost:27017"
	// log.Println("DOTENV...", username, password, clusterEndpoint)
	// logras = l["info"]
	// logras.Info("DOTENV...", username, password, clusterEndpoint)
	connectionURI := fmt.Sprintf(connectionStringTemplate, username, password, clusterEndpoint)
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
	if err != nil {
		log.Printf("Failed to create client: %v", err)
	}
	Ctx, Cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	err = client.Connect(Ctx)
	if err != nil {
		log.Printf("Failed to connect to cluster: %v", err)
	}
	// Force a connection to verify our connection string
	err = client.Ping(Ctx, nil)
	if err != nil {
		log.Fatal("Failed to ping cluster:", err)

	}

	log.Println("Connected to MongoDB!")
	return client, Ctx, Cancel
}

//ConnectDB2 ...
func (m *MongoStruct) ConnectDB2() (*mongo.Client, error) {
	username := os.Getenv("MONGODB_USERNAME")
	password := os.Getenv("MONGODB_PASSWORD")
	clusterEndpoint := os.Getenv("MONGODB_ENDPOINT")
	log.Println("DOTENV...", username, password, clusterEndpoint)
	connectionURI := fmt.Sprintf(connectionStringTemplate, username, password, clusterEndpoint)
	var err error
	m.MongoClient, err = mongo.NewClient(options.Client().ApplyURI(connectionURI))
	if err != nil {
		log.Printf("Failed to create client: %v", err)
		return nil, err
	}
	MONGOClient = m.MongoClient
	m.Ctx, m.Cancel = context.WithTimeout(context.Background(), connectTimeout*time.Second)
	err = m.MongoClient.Connect(m.Ctx)
	if err != nil {
		log.Printf("Failed to connect to cluster: %v", err)
		return nil, err
	}
	// Force a connection to verify our connection string
	err = m.MongoClient.Ping(m.Ctx, nil)
	if err != nil {
		log.Printf("Failed to ping cluster: %v", err)
		return nil, err
	}

	log.Println("Connected to MongoDB!")
	return m.MongoClient, nil
}

//RedisClient ...
var RedisClient *redis.Client

// func InitRedis(params ...string) {

// InitRedis ...
func InitRedis(params ...string) {

	var redisHost = os.Getenv("REDIS_HOST")
	var redisPassword = os.Getenv("REDIS_PASSWORD")

	db, _ := strconv.Atoi(params[0])

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPassword,
		DB:       db,
	})
	pong, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatal("Redis connection error...", err)
	}
	log.Println("Redis connections success...", pong)
}

// InitRedis2 ...
func InitRedis2() {
	var redisHost = os.Getenv("REDIS_HOST")
	var redisPassword = os.Getenv("REDIS_PASSWORD")
	params := [2]string{"1", "2"}
	db, _ := strconv.Atoi(params[0])

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPassword,
		DB:       db,
	})
	pong, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatal("Redis connection error...", err)
	}
	fmt.Println(pong, err)
}

//GetRedis ...
func GetRedis() *redis.Client {
	return RedisClient
}
