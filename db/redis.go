package db

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"

// 	_redis "github.com/go-redis/redis/v7"
// 	_ "github.com/lib/pq" //import postgres
// )

// //RedisClient ...
// var RedisClient *_redis.Client

// // InitRedis ...
// func InitRedis(params ...string) {

// 	var redisHost = os.Getenv("REDIS_HOST")
// 	var redisPassword = os.Getenv("REDIS_PASSWORD")

// 	db, _ := strconv.Atoi(params[0])

// 	RedisClient = _redis.NewClient(&_redis.Options{
// 		Addr:     redisHost,
// 		Password: redisPassword,
// 		DB:       db,
// 	})
// 	pong, err := RedisClient.Ping().Result()
// 	if err != nil {
// 		log.Fatal("Redis connection error...", err)
// 	}
// 	fmt.Println(pong, err)
// }

// //GetRedis ...
// func GetRedis() *_redis.Client {
// 	return RedisClient
// }
