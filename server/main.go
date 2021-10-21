package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/aryadiahmad4689/coba-protobuf/model/user"

	"google.golang.org/grpc"
)

//userService is a struct implementing UserServiceServer
type userService struct{}

//GreetUser return greeting message given the name and salutation
//in gRPC protocol
func (*userService) GetHello(ctx context.Context, req *user.GreetingRequest) (*user.GreetingResponse, error) {
	//business logic
	salutationMessage := fmt.Sprintf("Howdy, %s %s, nice to see you in the future!",
		req.Salutaion, req.Name)

	return &user.GreetingResponse{GreetingMessage: salutationMessage}, nil
}

func main() {
	//Create TCP Server on localhost:5001
	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("Failed to listen on port: %v", err)
	}
	//Create new gRPC server handler
	server := grpc.NewServer()

	//register gRPC UserService to gRPC server handler
	user.RegisterUserServiceServer(server, &userService{})

	//Run server
	if err := server.Serve(lis); err != nil {
		log.Fatal(err.Error())
	}
}
