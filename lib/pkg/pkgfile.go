package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const Filename = "package.json"

type PackageFile struct {
	Roms map[string]string `json:"roms"`
}

var PkgFile PackageFile

func LoadPackageFile() {
	if !exists(Filename) {
		fmt.Println("Rom hacks folder not initialized.")
	}

	rawJson, err := os.ReadFile(Filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = json.Unmarshal(rawJson, &PkgFile)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func InitPackageFile() {
	err := os.WriteFile(Filename, defaultPackage(), 0777)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func defaultPackage() []byte {
	defaultPackageJSON := PackageFile{
		Roms: map[string]string{},
	}
	data, _ := json.MarshalIndent(defaultPackageJSON, "", " ")
	return data
}

func exists(filepath string) bool {
	_, err := os.Stat(filepath)
	return !errors.Is(err, os.ErrNotExist)
}

func (p PackageFile) IsPresent(romhack string) bool {
	_, ok := p.Roms[romhack]
	return ok
}

func (p PackageFile) Write(romhack string, version string) error {
	p.Roms[romhack] = version
	data, _ := json.MarshalIndent(p, "", " ")
	return os.WriteFile(Filename, data, 0777)
}