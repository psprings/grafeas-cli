package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	pb "github.com/grafeas/grafeas/v1alpha1/proto"
	grafeasUtils "github.com/psprings/grafeas-cli/internal/grafeas"
	"github.com/psprings/grafeas-cli/internal/grafeas/projects"
	"github.com/spf13/cobra"
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
		conn, err := grafeasConfig.GrpcClient()
		// defer conn.Close()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("GRPC CONN: %v\n\n", conn)
		projectName = grafeasUtils.ProjectNameFormat(projectName)
		grafeasClient := projects.GrpcClient{Connection: conn}
		project := &pb.Project{Name: projectName}
		createProjectInput := &pb.CreateProjectRequest{Project: project}
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
		conn, err := grafeasConfig.GrpcClient()
		// defer conn.Close()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("GRPC CONN: %v\n\n", conn)
		pageSize = int32(size)
		grafeasClient := projects.GrpcClient{Connection: conn}
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
