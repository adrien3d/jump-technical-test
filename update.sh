#!/bin/bash
docker stop base-api
docker rm base-api
docker volume prune
docker rmi api-things-img:latest

docker build -t api-things-img .
docker run --name base-api -p 127.0.0.1:4000:4000 --link things-mongo:mongo -d api-things-img