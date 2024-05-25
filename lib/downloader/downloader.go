package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Souvikns/rpm/lib"
)

const rawContentUrl = "https://github.com/Souvikns/Rom-Hacks-Registry/raw/main/"

func Download(rom string, version string) {
	filename := lib.RegistryFile[rom].GetPatch(version)
	urlString := rawContentUrl + rom + "/" + filename
	fmt.Println(urlString)
	resp, err := http.Get(urlString)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	if err := os.WriteFile(filename, data, 0644); err != nil {
		panic(err)
	}
}
