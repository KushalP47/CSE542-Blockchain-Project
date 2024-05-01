package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addAccountCmd = &cobra.Command{
	Use:   "addAccount",
	Short: "Add a new account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Adding account...")
		// Implement your add account logic here
	},
}

func init() {
	rootCmd.AddCommand(addAccountCmd)
}
