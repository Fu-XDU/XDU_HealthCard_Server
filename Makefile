VERSION ?= 0.0.1
IMG_NAME ?= gin/golang-hello-server:${VERSION}

server:
	go build -o ./bin/server .

docker:
	docker build -t ${IMG_NAME} -f Dockerfile .

start:
	IMG_NAME=${IMG_NAME} docker-compose -f docker-compose.yaml up -d

stop:
	docker-compose -f docker-compose.yaml down

restart:
	docker-compose -f docker-compose.yaml restart