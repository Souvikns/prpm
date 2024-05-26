package cmd

import (
	"fmt"
	"os"

	"github.com/Souvikns/prpm/lib/pkg"
	"github.com/Souvikns/prpm/lib/registry"
	"github.com/spf13/cobra"
)

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download a rom hack from registry",
	Run: func(cmd *cobra.Command, args []string) {
		pkg.LoadPackageFile()
		registry.LoadRegistryFile()
		if len(args) != 0 {
			romHackName := args[0]
			ok := pkg.PkgFile.IsPresent(romHackName)
			if !ok {
				v, _ := registry.RegistryFile[romHackName].Latest()
				registry.Download(romHackName, v)
				pkg.PkgFile.Write(romHackName, v)
			}
		}
		for r, v := range pkg.PkgFile.Roms {
			fmt.Println("Downloading " + r + v)
			registry.Download(r, v)
		}
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
	downloadCmd.PersistentFlags().String("version", "", "Pass in a specific version to donload.")
}
