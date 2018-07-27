package cmd

import (
    "fmt"
    "os"

    homedir "github.com/mitchellh/go-homedir"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

var cfgFile string

// Grafeas vars
var grafeasScheme string
var grafeasHost string
var grafeasPort int
var grafeasGrpc bool
var grafeasBaseURL string

var RootCmd = &cobra.Command{
    Use:   "grafeas",
    Short: "A CLI to interact with a Grafeas server",
    Long: `This is a CLI to interact with a Grafeas server 
    and provides a way of hosting a simple API for interaction`,
}

func Execute() {
    if err := RootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func init() {
    cobra.OnInitialize(initConfig)

    RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.grafeas.yaml)")
    RootCmd.PersistentFlags().StringVar(&grafeasScheme, "scheme", "http", "The scheme of the Grafeas server")
    viper.BindPFlag("grafeas.scheme", RootCmd.PersistentFlags().Lookup("scheme"))
    RootCmd.PersistentFlags().StringVar(&grafeasHost, "host", "localhost", "The hostname of the Grafeas server")
    viper.BindPFlag("grafeas.host", RootCmd.PersistentFlags().Lookup("host"))
    RootCmd.PersistentFlags().IntVar(&grafeasPort, "port", 8080, "The port of the Grafeas server")
    viper.BindPFlag("grafeas.port", RootCmd.PersistentFlags().Lookup("port"))
    RootCmd.PersistentFlags().BoolVar(&grafeasGrpc, "grpc", false, "Whether to use GRPC client")
    viper.BindPFlag("grafeas.grpc", RootCmd.PersistentFlags().Lookup("grpc"))
    RootCmd.PersistentFlags().StringVar(&grafeasBaseURL, "base-url", "", "The base URL")
    viper.BindPFlag("grafeas.baseURL", RootCmd.PersistentFlags().Lookup("base-url"))
}

func initConfig() {
    if cfgFile != "" {
        viper.SetConfigFile(cfgFile)
    } else {
        // Find home directory.
        home, err := homedir.Dir()
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }

        viper.AddConfigPath(home)
        viper.SetConfigName(".grafeas.yaml")
    }

    viper.AutomaticEnv()

    if err := viper.ReadInConfig(); err == nil {
        fmt.Println("Using config file:", viper.ConfigFileUsed())
    }
}
