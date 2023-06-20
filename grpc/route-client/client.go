package main

import (
	"context"
	pb "github.com/NickYxy/GoTypingTest/grpc/route"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func runFirst(client pb.RouteGuideClient) {
	 feature, err := client.GetFeature(context.Background(), &pb.Point{
	 	Latitude: 310235000,
	 	Longitude: 121437403,
	 })

	 if err != nil {
		 log.Fatalln(err)
	 }

	 log.Println(feature)
}

func main() {
	conn, err := grpc.Dial("localhost:5000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalln("client cannot dial grpc server")
	}
	defer conn.Connect()

	client := pb.NewRouteGuideClient(conn)

	runFirst(client)
}