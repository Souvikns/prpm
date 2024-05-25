package lib

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDataStructure(t *testing.T) {
	raw := []byte(`{
    "pokemon-coral": {
        "base-game": "Pokemon - Crystal Version (UE) (V1.0) [C][!].gbc",
        "patch-versions": {
            "1.2": "coral2022demov1.2.ips"
        }
    }
}`)
	var registry map[string]RomHack

	err := json.Unmarshal(raw, &registry)
	t.Log(registry["pokemon-coral"].PatchVersions["1.2"])
	assert.Nil(t, err)

}

func TestLoadRegistry(t *testing.T) {
	LoadRegistryFile()
	t.Log(RegistryFile["pokemon-coral"])
	assert.NotNil(t, RegistryFile)
	assert.Equal(t, RegistryFile["pokemon-coral"].PatchVersions["1.2"], "coral2022demov1.2.ips")
}

func TestGetPatch(t *testing.T) {
	LoadRegistryFile()
	t.Log(RegistryFile["pokemon-coral"].GetPatch("1.2"))
	assert.Equal(t, RegistryFile["pokemon-coral"].GetPatch("1.2"), RegistryFile["pokemon-coral"].PatchVersions["1.2"])
}

func TestLatest(t *testing.T) {
	LoadRegistryFile()
	romHack := RegistryFile["pokemon-coral"]
	t.Log(romHack.Latest())
	version, _ := romHack.Latest()
	assert.Equal(t, version, "1.2")
}