package projects

import (
    "context"
    "fmt"
    "net/http"

    empty "github.com/golang/protobuf/ptypes/empty"
    grafeas "github.com/grafeas/client-go/v1alpha1"
    pb "github.com/grafeas/grafeas/v1alpha1/proto"
    "google.golang.org/grpc"
)

// Client :
type Client struct {
    Connection    *grpc.ClientConn
    GrafeasClient *grafeas.APIClient
    Context       context.Context
}

// GrpcClient :
type GrpcClient struct {
    Connection    *grpc.ClientConn
    GrafeasClient *grafeas.APIClient
    Context       context.Context
}

// ListProjects : wrapper around Grafeas
func (client *GrpcClient) ListProjects(listProjectsInput *pb.ListProjectsRequest) (*pb.ListProjectsResponse, error) {
    projectsClient := pb.NewGrafeasProjectsClient(client.Connection)
    fmt.Printf("GRPC CONN: %v\n", client.Connection)
    projResp, err := projectsClient.ListProjects(client.Context,
        listProjectsInput)
    return projResp, err
}


// ListProjects : wrapper around Grafeas for http
func (client *Client) ListProjects(listProjectsInput *grafeas.ListProjectsOpts) (*http.Response, error) {
    projectsClient := client.GrafeasClient.GrafeasProjectsApi
    _, projResp, err := projectsClient.ListProjects(client.Context,
        listProjectsInput)
    return projResp, err
}

// ListProjects : wrapper around Grafeas
// func (client *GrpcClient) ListProjects(listProjectsInput *pb.ListProjectsRequest) (*pb.ListProjectsResponse, error) {
// 	projectsClient := pb.NewGrafeasProjectsClient(client.Connection)
// 	fmt.Printf("GRPC CONN: %v\n", client.Connection)
// 	projResp, err := projectsClient.ListProjects(client.Context,
// 		listProjectsInput)
// 	return projResp, err
// }

// ListProjectsFormatted : wrapper around Grafeas
// func (client *GrpcClient) ListProjectsFormatted(listProjectsInput *pb.ListProjectsRequest) {
// 	projResp, err := client.ListProjects(listProjectsInput)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// }

// CreateProject : wrapper around Grafeas grpc
func (client *GrpcClient) CreateProject(createProjectInput *pb.CreateProjectRequest) (*empty.Empty, error) {
    projectsClient := pb.NewGrafeasProjectsClient(client.Connection)
    fmt.Printf("GRPC CONN: %v\n", client.Connection)
    projResp, err := projectsClient.CreateProject(client.Context,
        createProjectInput)
    return projResp, err
}

// CreateProject : wrapper around Grafeas for http
func (client *Client) CreateProject(createProjectInput grafeas.ApiProject) (*http.Response, error) {
    projectsClient := client.GrafeasClient.GrafeasProjectsApi
    _, projResp, err := projectsClient.CreateProject(client.Context,
        createProjectInput)
    return projResp, err
}
