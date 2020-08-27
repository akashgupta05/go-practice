package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/akashgupta05/go-practice/grpc/app"
	"google.golang.org/grpc"
)

//go:generate protoc --go_out=plugins=grpc:. ./app/app.proto

type Listener int

func (l *Listener) GetData(c context.Context, r *app.Request) (*app.Response, error) {
	reqData := r.Data
	fmt.Println("Receive", reqData)

	return &app.Response{Data: fmt.Sprintf("Nice : %s", reqData)}, nil
}

func main() {
	add, err := net.ResolveTCPAddr("tcp", "0.0.0.0:1234")
	if err != nil {
		log.Fatal(err)
	}

	inbound, err := net.ListenTCP("tcp", add)
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	listener := new(Listener)

	app.RegisterAppServer(server, listener)

	server.Serve(inbound)
	fmt.Println("Started")
}
