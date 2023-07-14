VERSION ?= 0.0.1
IMG_NAME ?= fuming/xdu_health_card_server:${VERSION}

server:
	go build -o ./bin/server .

docker:
	docker build -t ${IMG_NAME} -f Dockerfile .

start:
	IMG_NAME=${IMG_NAME} docker-compose -f docker-compose.yaml up -d

down:
	IMG_NAME=${IMG_NAME} docker-compose -f docker-compose.yaml down

restart:
	IMG_NAME=${IMG_NAME} docker-compose -f docker-compose.yaml restart

server-all:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./bin/server-darwin-amd64 .
	#CGO_ENABLED=0 GOOS=darwin GOARCH=386 go build -o ./bin/server-darwin-i386 .
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/server-linux-amd64 .
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o ./bin/server-linux-arm .
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o ./bin/server-linux-i386 .
	CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 go build -o ./bin/server-freebsd-amd64 .
	CGO_ENABLED=0 GOOS=freebsd GOARCH=386 go build -o ./bin/server-freebsd-i386 .
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./bin/server-windows-amd64 .
    CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o ./bin/server-windows-i386 .

