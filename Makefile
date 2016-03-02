all: protoc

protoc:
	protoc --go_out=plugins=grpc:. message_service.proto
