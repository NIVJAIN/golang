#!/bin/bash
# @sripal.jain@gmail.com
set -o errexit
# display_usage() {
#     echo
#     echo "USAGE: ./build_push_update_images.sh <version> [-h|--help] [--prefix=value] [--scan-images]"
#     echo "	version : Version of the sample app images (Required)"
#     echo "	-h|--help : Prints usage information"
#     echo "	--prefix: Use the value as the prefix for image names. By default, 'istio' is used"
#     echo -e "	--scan-images : Enable security vulnerability scans for docker images \n\t\t\trelated to bookinfo sample apps. By default, this feature \n\t\t\tis disabled."
#     exit 1
# }
# Check if docker is running
if ! docker info >/dev/null 2>&1; then
    echo "Docker does not seem to be running, run it first and retry"
    exit 1
fi
# display_usage

# if [ "$#" -ne 2 ]; then
#     echo "Incorrect parameters"
#     echo "Usage: master.sh <version> <prefix>"
#     exit 1
# fi

VERSION=$1
PREFIX=$2
MONGO=mongolang
REDIS=redisgolang
KAFKA_SWARM_CLUSTER_NAME=kafkacluster
declare -a CONTAINERS=($MONGO, $REDIS)

SCRIPTDIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
echo $SCRIPTDIR

if [ "$1" = "up" ]; then
    # pushd "$SCRIPTDIR/kafka-stack-docker-compose-master"
    # sh run.sh up
    # popd

    pushd "$SCRIPTDIR/rabbitmq-cluster_docker_compose/cluster_conf"
    sh run.sh up
    popd

    if [ "$( docker container inspect -f '{{.State.Status}}' $MONGO )" != "running" ]; then
        docker run -d --name $MONGO \
            -e MONGO_INITDB_ROOT_USERNAME=root \
            -e MONGO_INITDB_ROOT_PASSWORD=iloveblockchain \
            -p 27017:27017 \
            mongo
    fi
    if [ "$( docker container inspect -f '{{.State.Status}}' $REDIS )" != "running" ]; then
        docker run -itd --name $REDIS -p 6379:6379 redis
    fi
    sleep 5
    # run the golang application by executing below script
    sh bhag.sh
fi

if [ "$1" = "down" ]; then
    # container_name=reverent_mahavira
    pushd "$SCRIPTDIR/kafka-stack-docker-compose-master"
    sh run.sh down
    popd

    pushd "$SCRIPTDIR/rabbitmq-cluster_docker_compose/cluster_conf"
    sh run.sh down
    popd
  
    for i in "${CONTAINERS[@]}"
    do
        echo "$i"
        if [ "$( docker container inspect -f '{{.State.Status}}' $i )" == "running" ]; then
            echo "running"
            docker stop $i && docker rm $i
            # docker stop $(docker ps -q --filter ancestor=$container_name )
        fi
    done  
fi

if [ "$1" = "kafkaswarmu" ]; then
    pushd "$SCRIPTDIR/kafka-stack-docker-compose-master"
    docker swarm init
    docker stack deploy --compose-file zk-multiple-kafka-multiple.yml $KAFKA_SWARM_CLUSTER_NAME
    popd
fi
if [ "$1" = "kafkaswarmd" ]; then
    pushd "$SCRIPTDIR/kafka-stack-docker-compose-master"
    docker stack ls
    docker stack rm $KAFKA_SWARM_CLUSTER_NAME
    popd
fi
