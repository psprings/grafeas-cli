package cmd

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "io/ioutil"

    "github.com/spf13/viper"

    grafeas "github.com/grafeas/client-go/v1alpha1"
    pb "github.com/grafeas/grafeas/v1alpha1/proto"
    grafeasUtils "github.com/psprings/grafeas-cli/internal/grafeas"
    "github.com/psprings/grafeas-cli/internal/grafeas/projects"
    "github.com/spf13/cobra"
    "github.com/antihax/optional"
)

var projectName string
var createProjectCmd = &cobra.Command{
    Use:   "project",
    Short: "Creates a project",
    Long:  `Creates a Grafeas project`,
    Run: func(cmd *cobra.Command, args []string) {
        // fmt.Println("Creating project...")
        fmt.Printf("Creating project: %s\n", projectName)
        grafeasConfig := grafeasUtils.GenerateConfig()
        projectName = grafeasUtils.ProjectNameFormat(projectName)
        if viper.GetBool("grafeas.grpc") {
            conn, err := grafeasConfig.GrpcClient()
            defer conn.Close()
            if err != nil {
                log.Fatal(err)
            }
            fmt.Printf("GRPC CONN: %v\n\n", conn)
            grafeasClient := projects.GrpcClient{Connection: conn, Context: context.Background()}
            project := pb.Project{Name: projectName}
            createProjectInput := &pb.CreateProjectRequest{Project: &project}
            createProjectOutput, err := grafeasClient.CreateProject(createProjectInput)
            if err == nil {
                b, jsonErr := json.Marshal(createProjectOutput)
                if jsonErr != nil {
                    log.Fatal(jsonErr)
                }
                fmt.Println(b)
            } else {
                log.Fatal(err)
            }
        } else {
            newConfig := grafeas.NewConfiguration()
            newConfig.BasePath = grafeasConfig.BaseURL
            newClient := grafeas.NewAPIClient(newConfig)
            grafeasClient := projects.Client{Context: context.Background(), GrafeasClient: newClient}
            createProjectInput := grafeas.ApiProject{Name: projectName}
            createProjectOutput, err := grafeasClient.CreateProject(createProjectInput)
            if err == nil {
                b, jsonErr := json.Marshal(createProjectOutput)
                if jsonErr != nil {
                    log.Fatal(jsonErr)
                }
                fmt.Println(b)
            } else {
                log.Fatal(err)
            }
        }
    },
}

func createProjectCmdFlags() {
    createProjectCmd.Flags().StringVarP(&projectName, "name", "n", "", "The name of the project to be created. Can be 'foo' or long form 'projects/foo'")
}

var size int
var pageSize int32
var listProjectsCmd = &cobra.Command{
    Use:   "projects",
    Short: "Lists available projects",
    Long:  `Lists available projects`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Listing projects...")
        grafeasConfig := grafeasUtils.GenerateConfig()
        pageSize = int32(size)
        if viper.GetBool("grafeas.grpc") {
            conn, err := grafeasConfig.GrpcClient()
            // defer conn.Close()
            if err != nil {
                log.Fatal(err)
            }
            fmt.Printf("GRPC CONN: %v\n\n", conn)
            grafeasClient := projects.GrpcClient{Connection: conn, Context: context.Background()}
            listProjectsInput := &pb.ListProjectsRequest{PageSize: pageSize}
            listProjectsOutput, err := grafeasClient.ListProjects(listProjectsInput)
            if err == nil {
                b, jsonErr := json.Marshal(listProjectsOutput)
                if jsonErr != nil {
                    log.Fatal(jsonErr)
                }
                fmt.Println(b)
            } else {
                log.Fatal(err)
            }
        } else {
            newConfig := grafeas.NewConfiguration()
            newConfig.BasePath = grafeasConfig.BaseURL
            newClient := grafeas.NewAPIClient(newConfig)
            grafeasClient := projects.Client{Context: context.Background(), GrafeasClient: newClient}
            listProjectsInput := grafeas.ListProjectsOpts{PageSize: optional.NewInt32(pageSize)}
            listProjectsOutput, err := grafeasClient.ListProjects(&listProjectsInput)
            defer listProjectsOutput.Body.Close()
            bodyBytes, err2 := ioutil.ReadAll(listProjectsOutput.Body)
            if err2 != nil {
                log.Fatal(err2)
            }
            bodyString := string(bodyBytes)
            if err == nil {
                log.Println(bodyString)
            } else {
                log.Fatal(err)
            }
        }
    },
}

// ListProjectsInput :
type ListProjectsInput struct {
    Size int
}

func listProjectsCmdFlags() {
    listProjectsCmd.Flags().IntVarP(&size, "size", "s", 50, "The number of projects to return 'PageSize'")
}

func init() {
    // Create project
    createProjectCmdFlags()
    CreateCmd.AddCommand(createProjectCmd)
    // List projects
    listProjectsCmdFlags()
    ListCmd.AddCommand(listProjectsCmd)
}
