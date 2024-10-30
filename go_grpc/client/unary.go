package main

import (
	"context"
	pb "go_grpc/proto"
	"log"
	"time"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatal("Could not greet: %v", err)
	}
	log.Printf("%s", res.Message)
}
