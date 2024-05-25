package cmd

import (
	"fmt"
	"os"

	"github.com/Souvikns/rpm/lib"
	"github.com/Souvikns/rpm/lib/downloader"
	"github.com/spf13/cobra"
)

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download a rom hack from registry",
	Run: func(cmd *cobra.Command, args []string) {
		lib.LoadLockFile()
		lib.LoadRegistryFile()
		if len(args) != 0 {
			romHackName := args[0]
			fmt.Println(romHackName)
			ok := lib.LockFile.IsPresent(romHackName)
			if !ok {
				v, _ := lib.RegistryFile[romHackName].Latest()
				downloader.Download(romHackName, v)
				lib.LockFile.Write(romHackName, v)
			}
		}
		for r, v := range lib.LockFile.Roms {
			fmt.Println("Downloading " + r + v)
			downloader.Download(r, v)
		}
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
	downloadCmd.PersistentFlags().String("version", "", "Pass in a specific version to donload.")
}
