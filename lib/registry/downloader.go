package registry

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const rawContentUrl = "https://github.com/Souvikns/Rom-Hacks-Registry/raw/main/"
const registryUrl = "https://raw.githubusercontent.com/Souvikns/Rom-Hacks-Registry/main/registry.json"

func Download(rom string, version string) {
	filename, _ := RegistryFile[rom].GetPatch(version)
	urlString := rawContentUrl + rom + "/" + filename
	resp, err := http.Get(urlString)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := os.WriteFile(filename, data, 0644); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func LoadRegistryFile() {
	resp, err := http.Get(registryUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	if err := json.Unmarshal(data, &RegistryFile); err != nil {
		panic(err)
	}
}