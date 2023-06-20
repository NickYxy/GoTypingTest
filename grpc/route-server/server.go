package main

import (
	"context"
	"log"
	"net"

	pb "github.com/NickYxy/GoTypingTest/grpc/route"
	"google.golang.org/grpc"
)

type routeGuideServer struct {
	pb.UnimplementedRouteGuideServer
}

func (s *routeGuideServer) GetFeature(context.Context, *pb.Point) (*pb.Feature, error) {
	return nil, nil
}

func (s *routeGuideServer) ListFeatures(*pb.Rectangle, pb.RouteGuide_ListFeaturesServer) error {
	return nil
}

func (s *routeGuideServer) RecordRoute(pb.RouteGuide_RecordRouteServer) error {
	return nil
}

func (s *routeGuideServer) Recommend(pb.RouteGuide_RecommendServer) error {
	return nil
}

func newServer() *routeGuideServer {
	return &routeGuideServer{}
}

func main() {

	lis, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		log.Fatalln("cannot create a listener at the address")
	}

	grpcServer := grpc.NewServer()
	pb.RegisterRouteGuideServer(grpcServer, newServer())
	log.Fatalln(grpcServer.Serve(lis))
}
