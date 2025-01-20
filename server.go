package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/exec"
	"sync"

	pb "github.com/Surya-V-28/CI-CD-BuildServer/services/common/buildrequest"
	"github.com/go-git/go-git/v5"
	"google.golang.org/grpc"
)

// Server struct implements the BuildRequestServiceServer interface
type Server struct {
	pb.UnimplementedBuildRequestServiceServer
}

// helper function to run commands
func runCommand(dir, command string, args []string) error {
	cmd := exec.Command(command, args...)
	cmd.Dir = dir
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

// CreateBuildRequestMethod handles incoming gRPC requests
func (s *Server) CreateBuildRequestMethod(ctx context.Context, req *pb.CreateBuildRequest) (*pb.CreateBuildRespone, error) {
	log.Printf("Received request: GitUrl=%s, GitBranch=%s, CurrentUser=%s, Owner=%s, PublicRepo=%v, AppFrameWork=%s",
		req.GetGitUrl(), req.GetGitBranch(), req.GetCurrentUser(), req.GetOwner(), req.GetPublicRepo(), req.GetAppFrameWork())
	cloneDir := "/tmp/CiCD" // Default location for cloning repo
	// Ensure the target directory exists
	err := os.MkdirAll(cloneDir, os.ModePerm)
	if err != nil {
		log.Printf("Error creating clone directory: %v", err)
		return &pb.CreateBuildRespone{Status: "Failed to create clone directory"}, nil
	}

	// Clone the given repository to the specified directory
	log.Println("Cloning repository:", req.GetGitUrl())
	_, err = git.PlainClone(cloneDir, false, &git.CloneOptions{
		URL:      req.GetGitUrl(),
		Progress: os.Stdout,
	})
	if err != nil {
		log.Printf("Error cloning repository: %v", err)
		return &pb.CreateBuildRespone{Status: "Failed to clone repository"}, nil
	}

	// Run npm install in the cloned directory
	log.Println("Running 'npm install'...")
	if err := runCommand(cloneDir, "npm", []string{"install"}); err != nil {
		log.Printf("Error during npm install: %v", err)
		return &pb.CreateBuildRespone{Status: "Failed during npm install"}, nil
	}

	// Run npm build command
	log.Println("Running 'npm run build'...")
	if err := runCommand(cloneDir, "npm", []string{"run", "build"}); err != nil {
		log.Printf("Error during npm run build: %v", err)
		return &pb.CreateBuildRespone{Status: "Failed during npm run build"}, nil
	}
	// Respond back with success
	log.Println("Build completed successfully.")
	return &pb.CreateBuildRespone{Status: "Build Request Accepted"}, nil
}

func main() {
	// Use sync.WaitGroup to wait for the server and client to complete
	var wg sync.WaitGroup
	wg.Add(1)

	// Start the gRPC server in a goroutine
	go func() {
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
		wg.Done() // Signal that the server has started
	}()

	// Wait for the server to finish before exiting
	wg.Wait()
}
