PACKAGE_NAME=testing-golang

go-run:
	go run cmd/main.go

go-build:
	go build -o build/${PACKAGE_NAME} cmd/main.go