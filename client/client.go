package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"home.tech/proto/echo"

)

func main() {
	fmt.Printf("Hello")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := echo.NewEchoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.Echo(ctx, &echo.EchoRequest{Message: os.Args[1]})
	if err != nil {
		log.Fatalf("Could not echo: %v", err)
	}

	log.Printf("Response: %s", response.Message)
}
