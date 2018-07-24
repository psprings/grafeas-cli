package grafeas

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// Config :
type Config struct {
	Host             string
	Port             int
	ConnectionString string
}

// GenerateConfig :
func GenerateConfig() *Config {
	grafeasHost := viper.GetString("grafeas.host")
	grafeasPort := viper.GetInt("grafeas.port")
	connectionString := viper.GetString("grafeas.connectionString")
	if connectionString == "" {
		connectionString = fmt.Sprintf("%s:%d", grafeasHost, grafeasPort)
	}
	return &Config{
		Host:             grafeasHost,
		Port:             grafeasPort,
		ConnectionString: connectionString,
	}
}

// GrpcClient :
func (config *Config) GrpcClient() (*grpc.ClientConn, error) {
	connectionString := config.ConnectionString
	if connectionString == "" {
		connectionString = fmt.Sprintf("%s:%d", config.Host, config.Port)
	}
	conn, err := grpc.Dial(connectionString, grpc.WithInsecure())
	defer conn.Close()
	return conn, err
}

func projectPrefix() string {
	return "projects/"
}

// ProjectNameFormat : ensure that valid projectname is used
func ProjectNameFormat(projectName string) string {
	usePrefix := projectPrefix()
	if !strings.HasPrefix(projectName, usePrefix) {
		projectName = usePrefix + projectName
	}
	return projectName
}
