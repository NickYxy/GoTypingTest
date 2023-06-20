package main

import (
	"log"
	"net"

	pb "github.com/NickYxy/GoTypingTest/grpc/route"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		log.Fatalln("cannot create a listener at the address")
	}

	grpcServer = grpc.NewServer()
	pb.RegisterRouteGuideServer(grpc.Server)
}
