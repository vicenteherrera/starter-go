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
	if err := configure(); err != nil {
		fmt.Printf("%s \n\n", err)
		pflag.Usage()
		os.Exit(1)
	}

	log.SetOutput(os.Stdout)
	log.SetLevel(log.TraceLevel)
	log.Info("starter-go example program running")

	sample.ShowParams()

	client := analyzer.NewClient(viper.GetString("filename"))
	client.AnalyzeFile()

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

	pflag.StringP("filename", "f", "", "Name of the file to test")
	pflag.StringP("type", "t", "yaml", "File type")
	pflag.String("test", "privileged", "Test to perform on file")
	pflag.BoolP("break", "b", false, "Break on first error")

	pflag.VisitAll(func(flag *pflag.Flag) { viper.BindPFlag(flag.Name, flag) })
	pflag.Parse()

	if viper.Get("filename") == "" {
		return errors.New("filename is required")
	}

	return nil
}
