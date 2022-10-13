package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
	analyzer "github.com/vicenteherrera/starter-go/pkg/analyzer/containerfile"
	"github.com/vicenteherrera/starter-go/pkg/sample"
)

var cfgFile string

var version = "0.0.1"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "starter-go",
	Version: version,
	Short:   "A starter go program example",
	Long: `This is a starter go program example.
	
For more information on how to use it, execute:
starter-go help`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		if viper.GetString("filename") == "" {
			return errors.New("Filename parameter is required")
		}

		log.SetOutput(os.Stdout)
		log.SetLevel(log.TraceLevel)
		log.Info("starter-go example program running")

		// Main processing
		sample.ShowParams()
		client := analyzer.NewClient(viper.GetString("filename"))
		client.AnalyzeFile()

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.test.yaml)")

	// Cobra local flags, which will only run when this action is called directly.

	rootCmd.Flags().StringP("filename", "f", "", "Name of the file to test")
	rootCmd.Flags().StringP("type", "t", "yaml", "File type")
	rootCmd.Flags().String("test", "privileged", "Test to perform on file")
	rootCmd.Flags().BoolP("break", "b", false, "Break on first error")

	//rootCmd.MarkFlagRequired("filename")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".starter-go" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".starter-go")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

}
