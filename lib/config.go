package lib

import (
	"github.com/BurntSushi/toml"
)

// ConfigType type
type ConfigType struct {
	Version     string
	Environment string
	Database    struct {
		Driver   string
		Host     string
		Port     int
		Name     string
		User     string
		Password string
	} `toml:"database"`
	Jwt struct {
		Secret string
	} `toml:"jwt"`
	SSL struct {
		CertPath string
		KeyPath  string
	} `toml:"ssl"`
	Server struct {
		Port         int
		AllowOrigins []string
	} `toml:"server"`
	WebSocketServer struct {
		Port int
	} `toml:"webSocketServer"`
	Log struct {
		FileName         string
		NbFilesToArchive int
	} `toml:"log"`
	SQLLog struct {
		Level            uint
		Limit            float64
		DisplayOverLimit bool
	} `toml:"sql_log"`
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

// IsServerConfigCorrect checks if server config is correct
func IsServerConfigCorrect() bool {
	return len(Config.Server.AllowOrigins) > 0
}
