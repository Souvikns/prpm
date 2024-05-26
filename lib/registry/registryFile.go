package registry

import (
	"errors"

	"github.com/hashicorp/go-version"
)


type RomHack struct {
	BaseGame      string       `json:"base-game"`
	PatchVersions PatchVersion `json:"patch-versions"`
	System        string       `json:"system"`
}

type PatchVersion map[string]string

var RegistryFile map[string]RomHack


func (rh RomHack) GetPatch(version string) (string, error) {
	patch, ok := rh.PatchVersions[version]
	if !ok {
		return "", errors.New("version not available")
	}
	return patch, nil
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