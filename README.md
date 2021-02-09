### creating golang application
```
mkdir hello
cd hello
go mod init github.com/nivjain/hello
```

### Initial setup (Kafka & RabbitMQ via Docker)
```
MongoDB is required and Redis
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


