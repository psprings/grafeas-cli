package projects

import (
	"context"
	"fmt"

	empty "github.com/golang/protobuf/ptypes/empty"
	pb "github.com/grafeas/grafeas/v1alpha1/proto"
	"google.golang.org/grpc"
)

// GrpcClient :
type GrpcClient struct {
	Connection *grpc.ClientConn
}

// ListProjects : wrapper around Grafeas
func (client *GrpcClient) ListProjects(listProjectsInput *pb.ListProjectsRequest) (*pb.ListProjectsResponse, error) {
	projectsClient := pb.NewGrafeasProjectsClient(client.Connection)
	fmt.Printf("GRPC CONN: %v\n", client.Connection)
	projResp, err := projectsClient.ListProjects(context.Background(),
		listProjectsInput)
	return projResp, err
}

// ListProjectsFormatted : wrapper around Grafeas
// func (client *GrpcClient) ListProjectsFormatted(listProjectsInput *pb.ListProjectsRequest) {
// 	projResp, err := client.ListProjects(listProjectsInput)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// }

// CreateProject : wrapper around Grafeas
func (client *GrpcClient) CreateProject(createProjectInput *pb.CreateProjectRequest) (*empty.Empty, error) {
	projectsClient := pb.NewGrafeasProjectsClient(client.Connection)
	fmt.Printf("GRPC CONN: %v\n", client.Connection)
	projResp, err := projectsClient.CreateProject(context.Background(),
		createProjectInput)
	return projResp, err
}
