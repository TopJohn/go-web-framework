#!/bin/bash

function printHelp(){
    echo "Usage: "
    echo "  postgresql.sh <up|down>"
    echo "      up - start database."
    echo "      down - stop database."
}
MODE=$1

if [ $MODE == "up" ]; then
    echo "start postgres database."
    docker-compose up -d
elif [ $MODE == "down" ]; then
    echo "stop postgres database."
    docker-compose down
else
    printHelp
fi