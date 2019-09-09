package lib

import (
	"testing"

	"github.com/spf13/viper"
)

func TestInitConfig(t *testing.T) {
	InitConfig("../config.toml")

	timezone := viper.GetString("database.timzone") != "" && viper.GetString("database.timzone") != "Europe%2FParis"
	if timezone {
		t.Errorf("InitConfig - got: %t, want: %t.", timezone, false)
	}
}
