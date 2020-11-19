package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Process certificates and rotate if needed.",
	Args: func(cmd *cobra.Command, args []string) error {
		if authMode != "cli" && authMode != "env" {
			return fmt.Errorf("invalid auth mode '%s'. Expecting 'cli' or 'env'", authMode)
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Update...", configFile, authMode)
	},
}
