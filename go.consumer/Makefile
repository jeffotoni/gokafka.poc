# Makefile
.EXPORT_ALL_VARIABLES:

GO111MODULE=on
GOPROXY=direct
GOSUMDB=off

update:
	@echo "########## Compilando nossa API ... "
	rm -f go.*
	go mod init github.com/jeffotoni/go.consumer
	go build --trimpath -ldflags="-s -w" -o go.consumer main.go
	upx go.consumer
	@echo "buid completo..."
	@echo "\033[0;33m################ run #####################\033[0m"
	./go.consumer

build:
	@echo "########## Compilando nossa API ... "
	go build --trimpath -ldflags="-s -w" -o go.consumer main.go
	#upx go.consumer
	@echo "buid completo..."
	@echo "\033[0;33m################ run #####################\033[0m"
	./go.consumer

docker:
	@echo "########## Compilando nossa API ... "
	#CGO_ENABLED=0 
	go build --trimpath -ldflags="-s -w" -o go.consumer main.go
	upx go.consumer
	echo "-------------------------------------- Clean <none> images ---------------------------------------"
	#docker rmi $(docker images | grep "<none>" | awk '{print $3}') --force
	echo "\033[0;33m################################## build docker go.consumer ##################################\033[0m"
	docker build --no-cache -f Dockerfile -t jeffotoni/go.consumer .
	echo "\033[0;33m######################################### login aws and push ########################################\033[0m"
	echo "\033[0;32mGenerated\033[0m \033[0;33m[ok]\033[0m \n"
	#docker run -e ENV=DOCKER -e HOST=192.168.0.10:9092 -e DB_HOST=192.168.0.10 --rm --name go.consumer jeffotoni/go.consumer
	#docker ps -a | grep go.consumer
	
deploy.aws:
	@echo "########## Compilando nossa API ... "
	sh deploy.aws.sh
	@echo "fim"

deploy.google:
	@echo "########## Compilando nossa API ... "
	sh deploy.google.sh
	@echo "fim"