package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	authMode   string
	configFile string

	rootCmd = &cobra.Command{
		Use:   "az-acme",
		Short: "Manage TLS certificates in Azure",
		Long:  `Manage the creation of TLS certificates in Azure from Acme compliant CAs such as Let's Encrypt.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {

	updateCmd.Flags().StringVar(&configFile, "config", "", "config file to process")
	updateCmd.MarkFlagRequired("config")
	updateCmd.Flags().StringVar(&authMode, "auth", "cli", "authentication mode 'cli' or 'env'")

	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(initCmd)
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}
