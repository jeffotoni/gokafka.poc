# Makefile
.EXPORT_ALL_VARIABLES:

GO111MODULE=on
GOPROXY=direct
GOSUMDB=off

update:
	@echo "########## Compilando nossa API ... "
	rm -f go.*
	go mod init github.com/jeffotoni/server.game.blocks
	go build --trimpath -ldflags="-s -w" -o server.game.blocks main.go
	upx server.game.blocks
	@echo "buid completo..."
	@echo "\033[0;33m################ run #####################\033[0m"
	./server.game.blocks

build:
	@echo "########## Compilando nossa API ... "
	go build --trimpath -ldflags="-s -w" -o server.game.blocks main.go
	#upx server.game.blocks
	@echo "buid completo..."
	@echo "\033[0;33m################ run #####################\033[0m"
	./server.game.blocks

docker:
	@echo "########## Compilando nossa API ... "
	#CGO_ENABLED=0 
	go build --trimpath -ldflags="-s -w" -o server.game.blocks main.go
	upx server.game.blocks
	echo "-------------------------------------- Clean <none> images ---------------------------------------"
	#docker rmi $(docker images | grep "<none>" | awk '{print $3}') --force
	echo "\033[0;33m################################## build docker server.game.blocks ##################################\033[0m"
	docker build --no-cache -f Dockerfile -t jeffotoni/server.game.blocks .
	echo "\033[0;33m######################################### login aws and push ########################################\033[0m"
	echo "\033[0;32mGenerated\033[0m \033[0;33m[ok]\033[0m \n"
	#docker run -e ENV=DOCKER -e HOST=192.168.0.10:9092 -e DB_HOST=192.168.0.10 --rm --name server.game.blocks jeffotoni/server.game.blocks
	#docker ps -a | grep server.game.blocks
	
deploy.aws:
	@echo "########## Compilando nossa API ... "
	sh deploy.aws.sh
	@echo "fim"

deploy.google:
	@echo "########## Compilando nossa API ... "
	sh deploy.google.sh
	@echo "fim"