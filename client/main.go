package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aryadiahmad4689/coba-protobuf/model/user"

	"google.golang.org/grpc"
)

const userServiceAddress = "localhost:5001"

func main() {
	//Create connection to gRPC server
	conn, err := grpc.Dial(userServiceAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to service: %v", err)
		return
	}
	defer conn.Close()

	//Create new userService client
	userServiceClient := user.NewUserServiceClient(conn)

	//create connection timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := userServiceClient.GetHello(ctx, &user.GreetingRequest{
		Name:      "Marty Mc Fly",
		Salutaion: "Mr.",
	})
	if err != nil {
		log.Fatalf("Could not create request: %v", err)
	}

	//show response
	fmt.Println(response.GreetingMessage)
}
