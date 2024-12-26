package client

import (
	"context"
	"fmt"
	"log"

	pb "github.com/Surya-V-28/CI-CD-BuildServer/services/common/buildrequest"
	"google.golang.org/grpc"
)

func mainWelcome() {
	// Connect to the server at localhost:9000
	conn, err := grpc.Dial(":9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	// Create a new client for BuildRequestService
	client := pb.NewBuildRequestServiceClient(conn)

	// Prepare sample data for the CreateBuildRequest
	req := &pb.CreateBuildRequest{
		GitUrl:       "https://github.com/example/repository",
		GitBranch:    "main",
		CurrentUser:  "surya",
		Owner:        "surya-v-28",
		PublicRepo:   true,
		AppFrameWork: "Spring Boot", // Sample framework
	}

	// Call CreateBuildRequestMethod on the server
	resp, err := client.CreateBuildRequestMethod(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to call CreateBuildRequestMethod: %v", err)
	}

	// Print the response
	fmt.Printf("Response from server: %s\n", resp.GetStatus())
}
