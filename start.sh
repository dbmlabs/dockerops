#!/bin/bash
docker kill $(docker ps -aq)
docker rm -v $(docker ps -aq)
docker build -t 'dig' .
docker run -d -p 32794:9001 'dig'
docker rmi $(docker images -f "dangling=true" -q)
docker ps
