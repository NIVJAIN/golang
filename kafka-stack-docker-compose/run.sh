#!/bin/sh
set -e

if [ "$1" = "up" ];then
    echo "docker-compose up -d"
    docker-compose -f zk-multiple-kafka-multiple.yml up -d
fi


if [ "$1" = "down" ];then
    echo "docker-compose up -d"
    docker-compose -f zk-multiple-kafka-multiple.yml down
fi

if [ "$1" = "ps" ];then
    echo "docker-compose up -d"
    docker-compose -f zk-multiple-kafka-multiple.yml ps -a
fi

