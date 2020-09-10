#!/bin/bash
# Go Api server
# @jeffotoni
# 2020-06-02

echo "-------------------------------------- Clean <none> images ---------------------------------------"
docker rmi $(docker images | grep "<none>" | awk '{print $3}') --force

#echo "\033[0;33m################################## go build gokafka.poc ##################################\033[0m"
#GOOS=linux go build --trimpath -ldflags="-s -w" -o gokafka.poc main.go 
#upx gokafka.poc

echo "\033[0;33m################################## build docker gokafka.poc ##################################\033[0m"
docker build --no-cache -t gcr.io/projeto-eng1/gokafka.poc:latest .

docker login gcr.io

echo "\033[0;33m######################################### login aws and push ########################################\033[0m"
docker push gcr.io/projeto-eng1/gokafka.poc:latest
sleep 1
echo "\033[0;32mGenerated\033[0m \033[0;33m[ok apply]\033[0m \n"
#kubectl delete -f deployment.yaml
#@kubectl apply -f deployment.yaml
#echo "\033[0;32mGenerated\033[0m \033[0;33m[done]\033[0m \n"