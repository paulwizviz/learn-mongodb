#!/bin/bash

export MONGO_CMD_IMAGE=go-db/mongocmd:current
export MONGO_CMD_CONTAINER=mongoclient
export NETWORK=go-db_mongo-network

COMMAND="$1"

function build(){
    docker-compose -f ./build/docker-compose.yml build
}

function clean(){
    docker rmi -f ${MONGO_CMD_IMAGE}
    docker rmi -f $(docker images --filter "dangling=true" -q)
}

function shell(){
    docker exec -it ${MONGO_CMD_CONTAINER} /bin/sh
}

function start(){
    docker-compose -f ./deployment/docker-compose.yaml up
}

function stop(){
    docker-compose -f ./deployment/docker-compose.yaml down
}

case $COMMAND in
    "build")
        build
        ;;
    "clean")
        clean
        ;;
    "shell")
        shell
        ;;
    "start")
        start
        ;;
    "stop")
        stop
        ;;
    *)
        echo "$0 [ build | clean | shell | start | stop ]"
        ;;
esac