package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "github.com/bn222/grpctest/api"
)

var (
	port = flag.Int("port", 50051, "The server port")
)


type portServiceServer struct{
	pb.UnimplementedPortServiceServer
}

func (s *portServiceServer) CreatePort(ctx context.Context, req *pb.CreatePortRequest) (*pb.PortResponse, error) {
	details := req.Details
	fmt.Printf("Creating Port: VFID=%d, PFID=%d, VLAN=%d\n", details.Vfid, details.Pfid, details.Vlan)

	return &pb.PortResponse{Success: true}, nil
}

func (s *portServiceServer) DeletePort(ctx context.Context, req *pb.DeletePortRequest) (*pb.PortResponse, error) {
	details := req.Details
	fmt.Printf("Deleting Port: VFID=%d, PFID=%d, VLAN=%d\n", details.Vfid, details.Pfid, details.Vlan)

	return &pb.PortResponse{Success: true}, nil
}

func (s *portServiceServer) ClearAll(ctx context.Context, req *pb.ClearAllRequest) (*pb.ClearAllResponse, error) {
	fmt.Println("Clearing All Ports")
	return &pb.ClearAllResponse{Success: true}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterPortServiceServer(server, &portServiceServer{})
	log.Printf("server listening at %v", lis.Addr())

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
