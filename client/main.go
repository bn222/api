package main

import (
	"flag"
	"log"
	"fmt"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/bn222/grpctest/api"
)

 
var (
 	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewPortServiceClient(conn)

	createPortRequest := &pb.CreatePortRequest{
		Details: &pb.PortDetails{
			Vfid: 1,
			Pfid: 2,
			Vlan: 3,
		},
	}

	createPortResponse, err := client.CreatePort(context.Background(), createPortRequest)
	if err != nil {
		log.Fatalf("Error calling CreatePort: %v", err)
	}

	fmt.Printf("CreatePort Response: %v\n", createPortResponse)

	deletePortRequest := &pb.DeletePortRequest{
		Details: &pb.PortDetails{
			Vfid: 1,
			Pfid: 2,
			Vlan: 3,
		},
	}

	deletePortResponse, err := client.DeletePort(context.Background(), deletePortRequest)
	if err != nil {
		log.Fatalf("Error calling DeletePort: %v", err)
	}

	fmt.Printf("DeletePort Response: %v\n", deletePortResponse)

	clearAllRequest := &pb.ClearAllRequest{}

	clearAllResponse, err := client.ClearAll(context.Background(), clearAllRequest)
	if err != nil {
		log.Fatalf("Error calling ClearAll: %v", err)
	}

	fmt.Printf("ClearAll Response: %v\n", clearAllResponse)
}
