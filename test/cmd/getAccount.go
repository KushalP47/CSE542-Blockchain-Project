package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getAccountCmd = &cobra.Command{
	Use:   "getAccount",
	Short: "Retrieve an account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Retrieving account...")
		// Implement your get account logic here
	},
}

func init() {
	rootCmd.AddCommand(getAccountCmd)
}
