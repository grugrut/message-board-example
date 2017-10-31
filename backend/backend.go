package main

import (
	pb "github.com/grugrut/message-board-example"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()

}
