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

# display_usage

# if [ "$#" -ne 2 ]; then
#     echo "Incorrect parameters"
#     echo "Usage: master.sh <version> <prefix>"
#     exit 1
# fi

VERSION=$1
PREFIX=$2


SCRIPTDIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
echo $SCRIPTDIR


# pushd "$SCRIPTDIR/kafka-stack-docker-compose-master"
# # ls -rtl
# popd

# pushd "$SCRIPTDIR/rabbitmq-cluster_docker_compose/cluster_conf"
# # ls -rtl
# popd
#  docker run -d --name mongodb \
#       -e MONGO_INITDB_ROOT_USERNAME=root \
#       -e MONGO_INITDB_ROOT_PASSWORD=iloveblockchain \
#       -p 27017:27017 \
#       mongo

# docker run -itd -p 6379:6379 redis

# cd rabbitmq and run 

# cd kafka folder and run
