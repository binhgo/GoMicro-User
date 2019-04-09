build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/binhgo/GoMicro-User \
	proto/user/user.proto
	docker build -t user-service .

run:
	docker run --link=GoMicroPostgres:postgresDB --name GoMicroUserService -p 50053:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns -e DB_HOST=postgresDB -e DB_NAME=postgres -e DB_USER=postgres -e DB_PASSWORD=123456asd user-service