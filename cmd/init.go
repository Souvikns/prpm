package cmd

import (
	"github.com/Souvikns/prpm/lib/pkg"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize rm.json file to start downloading your rom hacks",
	Run: func(cmd *cobra.Command, args []string) {
		pkg.InitPackageFile()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
