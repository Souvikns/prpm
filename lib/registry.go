package lib

import (
	"encoding/json"
	"github.com/hashicorp/go-version"
	"io"
	"net/http"
)

const registryUrl = "https://raw.githubusercontent.com/Souvikns/Rom-Hacks-Registry/main/registry.json"

type RomHack struct {
	BaseGame      string       `json:"base-game"`
	PatchVersions PatchVersion `json:"patch-versions"`
}

type PatchVersion map[string]string

var RegistryFile map[string]RomHack

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
	if err := json.Unmarshal(data, &RegistryFile); err != nil {
		panic(err)
	}
}

func (rh RomHack) GetPatch(version string) string {
	patch, ok := rh.PatchVersions[version]
	if !ok {
		panic("Version not available")
	}
	return patch
}

func (rh RomHack) Latest() (string, string) {
	latest, _ := version.NewVersion("0.0.0")
	for key := range rh.PatchVersions {
		v, _ := version.NewVersion(key)
		if latest.LessThanOrEqual(v) {
			latest = v
		}
	}
	return latest.Original(), rh.PatchVersions[latest.Original()]
}
