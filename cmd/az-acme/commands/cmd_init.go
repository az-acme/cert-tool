package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise empty configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Initialise...")
	},
}
