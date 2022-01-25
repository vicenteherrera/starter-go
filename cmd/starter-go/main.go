package main

import (
	"errors"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	if err := configure(); err != nil {
		fmt.Printf("%s \n\n", err)

		pflag.Usage()
		os.Exit(1)
	}

	log.SetOutput(os.Stdout)
	log.SetLevel(log.TraceLevel)
	log.Info("Started starter-go ")
	log.Info("Server name: " + viper.GetString("server_name"))
	log.Info("Server host: " + viper.GetString("host"))
	log.Info("Server port: " + viper.GetString("port"))

	//log.Fatal(loop...)
}

func configure() error {
	viper.AutomaticEnv()

	viper.SetConfigName("config")            // name of config file (without extension)
	viper.SetConfigType("yaml")              // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("$HOME/.starter-go") // call multiple times to add many search paths
	viper.AddConfigPath(".")                 // optionally look for config in the working directory
	err := viper.ReadInConfig()              // Find and read the config file
	if err != nil {                          // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	pflag.String("server_name", "", "Server name")
	pflag.Bool("test", false, "Run in test mode")
	pflag.String("host", "localhost", "Host to use")
	pflag.Int("port", 8080, "Port to use")

	pflag.VisitAll(func(flag *pflag.Flag) { viper.BindPFlag(flag.Name, flag) })
	pflag.Parse()

	if viper.Get("server_name") == "" {
		return errors.New("server_token is required")
	}

	return nil
}
