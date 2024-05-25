package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const Filename = "roms.json"

type PackageFile struct {
	Roms map[string]string `json:"roms"`
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

var LockFile PackageFile

func LoadLockFile() {
	if !exists(Filename) {
		fmt.Println("Rom hacks folder not initialized.")
	}

	rawJson, err := os.ReadFile(Filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = json.Unmarshal(rawJson, &LockFile)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func InitializeLockFile() error {
	err := os.WriteFile(Filename, getInitLock(), 0777)
	if err != nil {
		return err
	}
	return nil
}

func getInitLock() []byte {
	newLock := PackageFile{
		Roms: map[string]string{},
	}

	data, _ := json.MarshalIndent(newLock, "", " ")

	return data
}

func exists(filepath string) bool {
	_, err := os.Stat(filepath)
	return !errors.Is(err, os.ErrNotExist)
}
