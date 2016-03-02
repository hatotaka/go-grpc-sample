package main

import (
	"golang.org/x/net/context"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"

	ggs "github.com/hatotaka/go-grpc-sample"
)

type messageService struct {
	messages []*ggs.Message
	m        sync.Mutex
	msg      chan *ggs.Message
}

func NewMessageService() *messageService {
	c := make(chan *ggs.Message, 128)
	return &messageService{msg: c}
}

func (this *messageService) WatchMessage(p *ggs.WatchType, stream ggs.MessageService_WatchMessageServer) error {

	for {
		stream.Send(<-this.msg)
	}

	return nil
}

func (this *messageService) SendMessage(c context.Context, m *ggs.Message) (*ggs.Result, error) {

	this.msg <- m

	return new(ggs.Result), nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()

	ggs.RegisterMessageServiceServer(server, NewMessageService())
	server.Serve(lis)
}
