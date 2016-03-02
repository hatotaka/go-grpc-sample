package main

import (
	//	"fmt"
	"io"
	//	"strconv"
	"log"

	ggs "github.com/hatotaka/go-grpc-sample"
	"github.com/mattn/sc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func send(body string) error {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return err
	}
	defer conn.Close()
	client := ggs.NewMessageServiceClient(conn)

	message := &ggs.Message{
		Body: body,
	}
	_, err = client.SendMessage(context.Background(), message)
	if err != nil {
		log.Println(err)
	}
	return err
}

func watch() error {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return err
	}
	defer conn.Close()
	client := ggs.NewMessageServiceClient(conn)

	stream, err := client.WatchMessage(context.Background(), new(ggs.WatchType))
	if err != nil {
		log.Println(err)
		return err
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
			return err
		}

		log.Println(message.Body)
	}
	return nil
}

func main() {
	(&sc.Cmds{
		{
			Name: "watch",
			Desc: "watch: watch message",
			Run: func(c *sc.C, args []string) error {
				return watch()
			},
		},
		{
			Name: "send",
			Desc: "send [message]: send message",
			Run: func(c *sc.C, args []string) error {
				if len(args) != 1 {
					return sc.UsageError
				}
				messageBody := args[0]
				return send(messageBody)
			},
		},
	}).Run(&sc.C{})
}
