package lib

import (
	"github.com/BurntSushi/toml"
)

// ConfigType type
type ConfigType struct {
	Version     string
	Environment string
	Database    databaseType `toml:"database"`
	Jwt         jwtType      `toml:"jwt"`
	SqlLog      sqlLogType   `toml:"sql_log"`
}

type databaseType struct {
	Driver   string
	Host     string
	Port     int
	Name     string
	User     string
	Password string
}

type jwtType struct {
	Secret string
}

type sqlLogType struct {
	Level uint
	Limit uint
}

// Config variable
var Config ConfigType

// InitConfig : Lecture du fichier de configuration
func InitConfig() {
	// Lecture du fichier de configuration
	// -----------------------------------
	if _, err := toml.DecodeFile("config.toml", &Config); err != nil {
		CheckError(err, -1)
	}
}

// IsDatabaseConfigCorrect : La configuration de la base de données est-elle correcte ?
func IsDatabaseConfigCorrect() bool {
	return Config.Database.Driver != "" && Config.Database.Name != "" && Config.Database.Host != ""
}
