package main

import (
	pb "buildingapp/services/common/buildrequest"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

// Server struct implements the BuildRequestServiceServer interface
type Server struct {
	pb.UnimplementedBuildRequestServiceServer
}

// CreateBuildRequestMethod handles incoming gRPC requests
func (s *Server) CreateBuildRequestMethod(ctx context.Context, req *pb.CreateBuildRequest) (*pb.CreateBuildRespone, error) {
	log.Printf("Received request: GitUrl=%s, GitBranch=%s, CurrentUser=%s, Owner=%s, PublicRepo=%v, AppFrameWork=%s",
		req.GetGitUrl(), req.GetGitBranch(), req.GetCurrentUser(), req.GetOwner(), req.GetPublicRepo(), req.GetAppFrameWork())

	// Create a response
	resp := &pb.CreateBuildRespone{
		Status: "Build Request Accepted",
	}
	return resp, nil
}

func main() {

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	// Create a gRPC server
	grpcServer := grpc.NewServer()

	// Register the BuildRequestServiceServer
	pb.RegisterBuildRequestServiceServer(grpcServer, &Server{})

	log.Println("Server is listening on port 9000...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
