### creating golang application
```
mkdir hello
cd hello
go mod init github.com/nivjain/hello
```

### Initial setup (Kafka & RabbitMQ via Docker)
```
MongoDB is required and Redis
 docker run -d --name mongodb \
      -e MONGO_INITDB_ROOT_USERNAME=root \
      -e MONGO_INITDB_ROOT_PASSWORD=iloveblockchain \
      -p 27017:27017 \
      mongo
 docker run -itd -p 6379:6379 redis
cd kafka-stack-docker-compose-master
sh run.sh up or docker-compose -f zk-multiple-kafka-multiple.yml up -d

cd <pgm dir>
cd rabbitmq-cluster_docker_compose
cd cluster_conf
cd docker-compose up -d
```

### how to run this app
```
1. git clone url
2. cd <clone url>
3. sh bhag.sh or go build && ./7-ginInterfaceMongoDBRabbitMQ-JWT-RTL
4. check file name LOGIN.rest and check all the options in that file. check all the files that ends with extension ".rest"
```

### How to run kafka consumers and rabbitmq consumers

## RabbitMQ
```
RABBIT.REST
click send request POST {{baseUrl}}/rabbit/push
cd 9-rabbitmq-con
go run consumer.go

```

## Kafka
```
KAFKA.
click send request to POST {{baseUrl}}/kafka/push
cd /gofka-master/conzumer/src
go run conjumer.go
```
