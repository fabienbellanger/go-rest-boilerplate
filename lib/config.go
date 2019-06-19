package lib

import (
	"strings"

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
		Timezone string
		Charset  string
		Engine   string
	} `toml:"database"`
	Jwt struct {
		Secret string
	} `toml:"jwt"`
	ORM struct {
		EnabledLogs           bool
		MaxIdleConnections    int
		MaxOpenConnections    int
		MaxLifetimeConnection int
	} `toml:"orm"`
	SSL struct {
		CertPath string
		KeyPath  string
	} `toml:"ssl"`
	Server struct {
		Port            int
		ReadTimeout     int
		WriteTimeout    int
		ShutdownTimeout int
		AllowOrigins    []string
	} `toml:"server"`
	WebSocketServer struct {
		Port int
	} `toml:"webSocketServer"`
	Log struct {
		DirPath          string
		FileName         string
		NbFilesToArchive int
		EnableAccessLog  bool
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
func InitConfig(file string) {
	// Lecture du fichier de configuration
	// -----------------------------------
	_, err := toml.DecodeFile(file, &Config)
	CheckError(err, 1)

	// On converti le / du timezone de la base de donnée en %2F
	// --------------------------------------------------------
	Config.Database.Timezone = strings.Replace(Config.Database.Timezone, "/", "%2F", -1)
}

// IsDatabaseConfigCorrect : La configuration de la base de données est-elle correcte ?
func IsDatabaseConfigCorrect() bool {
	return Config.Database.Driver != "" && Config.Database.Name != "" && Config.Database.Host != ""
}

// IsServerConfigCorrect checks if server config is correct
func IsServerConfigCorrect() bool {
	return len(Config.Server.AllowOrigins) > 0
}
