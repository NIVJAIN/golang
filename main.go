package main

import (
	"fmt"

	"github.com/nivjain/7-ginInterfaceMongoDBRabbitMQ-JWT-RTL/kafka"

	// "log"
	"net/http"
	"os"
	"runtime"

	"github.com/sirupsen/logrus"

	"github.com/nivjain/7-ginInterfaceMongoDBRabbitMQ-JWT-RTL/db"
	"github.com/nivjain/7-ginInterfaceMongoDBRabbitMQ-JWT-RTL/middlewares/logger"
	"github.com/nivjain/7-ginInterfaceMongoDBRabbitMQ-JWT-RTL/middlewares/recover"
	"github.com/nivjain/7-ginInterfaceMongoDBRabbitMQ-JWT-RTL/repositories"

	"github.com/nivjain/7-ginInterfaceMongoDBRabbitMQ-JWT-RTL/rabbitmq"
	"github.com/nivjain/7-ginInterfaceMongoDBRabbitMQ-JWT-RTL/routes"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/twinj/uuid"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	// swagger embed files
	// gin-swagger middleware
	// "github.com/swaggo/files" // swagger embed files
	// "github.com/swaggo/gin-swagger" // gin-swagger middleware

	ginSwagger "github.com/swaggo/gin-swagger"
	// gin-swagger middleware
)

// CollexionPool ...
var CollexionPool = make(map[string]*mongo.Collection)

//Clxpool ...
type Clxpool map[string]*mongo.Collection

//Init ...
func init() {
	var logams = new(logger.Logams)
	logsetup := logams.SetLogger("info")
	logpool["info"] = logsetup
	log = logsetup
	// log.Println("Server.go init fun called....")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file, please create one in the root directory")
	}

}

var logpool = make(map[string]*logrus.Logger)
var log *logrus.Logger

func main() {
	port := os.Getenv("PORT")
	log.Println("dddd")
	// log.Info("debug")
	// log.Debug("Debugams")
	// log.Error("erororo")
	// log.Errorln("errorline")

	r := gin.Default()
	r.Use(logger.LoggerToFile(), recover.Recover())
	r.Use(CORSMiddleware())
	r.Use(RequestIDMiddleware())
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(gin.Recovery())
	// MongoDb & Redis Database connections ...
	dba, _, _ := db.ConnectDB(logpool)
	// log.Println("DB::>>", dba)
	//Start Redis on database 1 - it's used to store the JWT but you can use it for anythig else
	//Example: db.GetRedis().Set(KEY, VALUE, at.Sub(now)).Err()
	db.InitRedis("1")

	// Kafka setup
	kafkaConnection := new(kafka.Connection)
	kafkainit, err := kafkaConnection.KafkaInitConnection("senz", log)
	if err != nil {
		log.Fatal("KafkaInitError", err)
	}
	// err = kafkainit.KafkaPublish("Hello Rama hello krishna")
	// if err != nil {
	// 	log.Info("Kafka Publish Error", err)
	// }

	rabbitConn := GetRabbitMQConnection(logpool)
	collections := GetCollectionPool(dba)

	// Routes profile
	var userRouter = new(routes.UserRouter)
	var rabRouter = new(routes.RabbitRouter)
	var kafRouter = new(routes.KafkaRouter)
	// HTML rendering ...
	r.LoadHTMLGlob("./public/html/*")
	r.Static("/public", "./public")
	r.Static("/datadog", "./datadog")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"ginBoilerplateVersion": "v0.03",
			"goVersion":             runtime.Version(),
		})
	})

	// Repo initialization for db and rabbitmq setup and then Routes setup ...
	mClientRepo := repositories.SetMongoClient(dba, collections, rabbitConn, logpool, kafkainit)
	rabRouter.Routes(dba, collections, rabbitConn, logpool, mClientRepo, r)
	userRouter.Routes(dba, collections, logpool, mClientRepo, r)
	kafRouter.Routes(mClientRepo, r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Default route if any of the above routes doesnt match ...
	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{"message": "Oops page not found!!!"})
	})
	log.Fatal(r.Run(":" + port))

}

//GetCollectionPool ...
func GetCollectionPool(db *mongo.Client) map[string]*mongo.Collection {
	users := db.Database("test").Collection("users")
	people := db.Database("test").Collection("people")
	tokens := db.Database("test").Collection("tokens")

	CollexionPool["users"] = users
	CollexionPool["people"] = people
	CollexionPool["tokens"] = tokens

	return CollexionPool
}

//GetCollectionPool2 ...
func GetCollectionPool2(db *mongo.Client) *Clxpool {
	users := db.Database("test").Collection("users")
	people := db.Database("test").Collection("people")
	kp := make(Clxpool)
	kp["users"] = users
	kp["people"] = people
	return &kp
}

// GetRabbitMQConnection ...
func GetRabbitMQConnection(logpool map[string]*logrus.Logger) *rabbitmq.Connection {
	conn := rabbitmq.NewConnection("my-producer", "my-exchange", []string{"queue-1"}, logpool)
	if err := conn.Connect(); err != nil {
		log.Error(err)
		panic(err)
	}
	if err := conn.BindQueue(); err != nil {
		log.Error(err)
		panic(err)
	}
	log.Println("RabbitMQ connection success...")
	return conn
}

//CORSMiddleware ...
//CORS (Cross-Origin Resource Sharing)
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

//RequestIDMiddleware ...
//Generate a unique ID and attach it to each request for future reference or use
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := uuid.NewV4()
		c.Writer.Header().Set("X-Request-Id", uuid.String())
		c.Next()
	}
}
