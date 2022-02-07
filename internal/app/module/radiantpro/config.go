package radiantpro

import (
	"io/ioutil"

	"github.com/W3-Partha/Radical/internal/pkg/utils"
	radicalLogger "github.com/W3-Partha/Radical/logger"
)

var CompareExcept = []string{"@RadicalGenerateTime"}

func (c *Container) GenConfig() {
	if utils.IsExist(c.RadiantProFile) {
		radicalLogger.Log.Fatalf("radiant pro toml exist")
		return
	}

	err := ioutil.WriteFile("radiantpro.toml", []byte(RadiantToml), 0644)
	if err != nil {
		radicalLogger.Log.Fatalf("write radiant pro toml err: %s", err)
		return
	}
}
