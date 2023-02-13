#!/bin/bash
docker stop jump-technical-test
docker rm jump-technical-test
docker volume prune
docker rmi api-things-img:latest

docker build -t api-things-img .
docker run --name jump-technical-test -p 127.0.0.1:4000:4000 --link things-mongo:mongo -d api-things-img