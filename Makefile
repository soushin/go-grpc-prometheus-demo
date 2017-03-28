
generate-proto:
	@echo "--> Generating Go files"
	protoc -I protobuf/ --go_out=plugins=grpc:protobuf/ protobuf/echo.proto
	@echo ""

start:
	@echo "--> Create and start containers"
	GOOS=linux GOARCH=386 go build -o ./server/server ./server/server.go
	GOOS=linux GOARCH=386 go build -o ./client/client ./client/client.go
	docker-compose up -d
	rm -rf ./server/server
	rm -rf ./client/client
	@echo ""

stop:
	@echo "--> Stop containers"
	docker-compose down
	@echo ""
