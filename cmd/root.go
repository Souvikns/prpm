package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "rhc",
	Short: "A CLI program to manage romhacks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Flags().GetString("name"))
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}