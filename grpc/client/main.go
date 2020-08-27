package main

import (
	"bufio"
	"context"
	"log"
	"os"

	"github.com/akashgupta05/go-practice/grpc/app"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := app.NewAppClient(conn)
	in := bufio.NewReader(os.Stdin)
	for {
		line, _, err := in.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		reply, err := c.GetData(context.Background(), &app.Request{Data: string(line)})
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Response: %v", reply)
	}
}
