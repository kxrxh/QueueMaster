package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use: "client command",
}

// init is called after the command line has been parsed and is used to
// initialize the viper configuration.
//
// It sets up the flags for the submit command and marks the task type and
// payload flags as required.
func init() {
	cobra.OnInitialize(initConfig)

	submitCmd.Flags().StringVarP(&taskType, "type", "T", "", "task type")
	submitCmd.Flags().StringVarP(&taskPayload, "payload", "P", "", "task payload")

	submitCmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		if taskType == "" || taskPayload == "" {
			return fmt.Errorf("Both task type and task payload are required")
		}
		return nil
	}

	submitCmd.MarkFlagRequired("type")
	submitCmd.MarkFlagRequired("payload")

	rootCmd.AddCommand(submitCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetConfigName("client_config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		os.Exit(1)
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// It is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
