package cmd

import (
	"fmt"
	"os"

	"github.com/Souvikns/rpm/lib"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize rm.json file to start downloading your rom hacks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		err := lib.InitializeLockFile()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
