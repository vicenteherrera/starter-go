package main

import (
	"errors"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	analyzer "github.com/vicenteherrera/starter-go/pkg/analyzer/containerfile"
	"github.com/vicenteherrera/starter-go/pkg/sample"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {

	// Check for errors in configuration
	if err := configure(); err != nil {
		fmt.Printf("%s \n\n", err)
		pflag.Usage()
		os.Exit(1)
	}

	// Start log
	log.SetOutput(os.Stdout)
	log.SetLevel(log.TraceLevel)
	log.Info("starter-go example program running")

	// Main processing
	sample.ShowParams()
	client := analyzer.NewClient(viper.GetString("filename"))
	client.AnalyzeFile()

}

func configure() error {

	// Activate getting env variables with viper.Get
	viper.AutomaticEnv()

	// Config file setup
	viper.SetConfigName("config")            // name of config file (without extension)
	viper.SetConfigType("yaml")              // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("$HOME/.starter-go") // call multiple times to add many search paths
	viper.AddConfigPath(".")                 // optionally look for config in the working directory
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignored error
		} else {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}

	// Command line parameter setup (precedence over config file)
	pflag.StringP("filename", "f", "", "Name of the file to test")
	pflag.StringP("type", "t", "yaml", "File type")
	pflag.String("test", "privileged", "Test to perform on file")
	pflag.BoolP("break", "b", false, "Break on first error")

	// Bind all flags for viper.Get
	pflag.VisitAll(func(flag *pflag.Flag) { viper.BindPFlag(flag.Name, flag) })

	// Parse configuration
	pflag.Parse()

	// Validation of parameters
	if viper.Get("filename") == "" {
		return errors.New("filename is required")
	}

	return nil
}
