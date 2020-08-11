PACKAGE_NAME:=hello-world-go
PACKAGE_VERSION=0.0.1
GIT_DIR=$(shell pwd)
DOCKER_HUB=pablon27

 
docker-push:
	docker build -t ${PACKAGE_NAME}:${PACKAGE_VERSION} .
	docker tag ${PACKAGE_NAME}:${PACKAGE_VERSION} ${DOCKER_HUB}/${PACKAGE_NAME}:${PACKAGE_VERSION} 
	docker push ${DOCKER_HUB}/${PACKAGE_NAME}:${PACKAGE_VERSION}

docker-run:
	docker run -it --rm ${PACKAGE_NAME}:${PACKAGE_VERSION}

docker-build:
	docker build -t ${PACKAGE_NAME}:${PACKAGE_VERSION} .

go-run-build:
	./build/${PACKAGE_NAME}

go-build:
	go build -o build/${PACKAGE_NAME} main.go 

go-run:
	go run main.go

go-test:
	go test -v

go-cover:
	go test -cover