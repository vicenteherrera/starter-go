package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
	analyzer "github.com/vicenteherrera/starter-go/pkg/analyzer/containerfile"
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

		filename := viper.GetViper().GetString("filename")

		if filename == "" {
			return errors.New("filename parameter is required")
		}

		log.SetOutput(os.Stdout)
		log.SetLevel(log.TraceLevel)
		log.Info("starter-go example program running")

		// Main processing
		client := analyzer.NewClient(filename)
		client.AnalyzeFile()

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.starter-go.yaml)")

	// Cobra local flags, which will only run when this action is called directly.

	rootCmd.Flags().StringP("filename", "f", "", "Name of the file to test")
	rootCmd.Flags().StringP("type", "t", "yaml", "File type")
	rootCmd.Flags().String("test", "privileged", "Test to perform on file")
	rootCmd.Flags().BoolP("break", "b", false, "Break on first error")

	viper.BindPFlag("config", rootCmd.Flags().Lookup("config"))
	viper.BindPFlag("filename", rootCmd.Flags().Lookup("filename"))
	viper.BindPFlag("type", rootCmd.Flags().Lookup("type"))
	viper.BindPFlag("test", rootCmd.Flags().Lookup("test"))
	viper.BindPFlag("break", rootCmd.Flags().Lookup("break"))

	viper.SetDefault("type", "yaml")
	viper.SetDefault("test", "yaml")

	// rootCmd.MarkFlagRequired("filename")
	// rootCmd.MarkFlagsRequiredTogether("username", "password")
	// rootCmd.MarkFlagsMutuallyExclusive("json", "yaml")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		if _, err := os.Stat(cfgFile); errors.Is(err, os.ErrNotExist) {
			fmt.Fprintln(os.Stderr, "Config file not found:", cfgFile)
		} else {
			viper.SetConfigFile(cfgFile)
		}
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
