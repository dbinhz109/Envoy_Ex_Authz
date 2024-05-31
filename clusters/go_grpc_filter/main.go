package main

import (
	"context"
	"fmt"
	"net"

	auth_pb "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	"google.golang.org/grpc"
)

type AuthServer struct{}

func (server *AuthServer) Check(ctx context.Context, request *auth_pb.CheckRequest) (*auth_pb.CheckResponse, error) {
	// Extract path and headers
	path := request.Attributes.Request.Http.Path[1:]
	headers := request.Attributes.Request.Http.Headers
	fmt.Println("Path:", path)

	// Check if path is /private
	if path == "private" {
		// Check for the x-allow-private header
		if val, ok := headers["x-allow-private"]; ok && val == "admin" {
			fmt.Println("allowed private request with x-allow-private=admin")
		} else {
			fmt.Println("blocked private request")
			return nil, fmt.Errorf("❌ private request not allowed")
		}
	}

	return &auth_pb.CheckResponse{
		HttpResponse: &auth_pb.CheckResponse_OkResponse{
			OkResponse: &auth_pb.OkHttpResponse{},
		},
	}, nil
}

func main() {
	// struct with check method
	endPoint := fmt.Sprintf("localhost:%d", 3001)
	listen, err := net.Listen("tcp", endPoint)
	if err != nil {
		fmt.Printf("Failed to listen on %s: %v\n", endPoint, err)
		return
	}

	grpcServer := grpc.NewServer()
	// register envoy proto server
	server := &AuthServer{}
	auth_pb.RegisterAuthorizationServer(grpcServer, server)

	fmt.Println("Server started at port 3001")
	if err := grpcServer.Serve(listen); err != nil {
		fmt.Printf("Failed to serve gRPC server: %v\n", err)
	}
}
