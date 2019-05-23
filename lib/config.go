package lib

import (
	"github.com/BurntSushi/toml"
)

// ConfigType type
type ConfigType struct {
	Version         string
	Environment     string
	Database        databaseType    `toml:"database"`
	Jwt             jwtType         `toml:"jwt"`
	Log             logType         `toml:"log"`
	Server          server          `toml:"server"`
	WebSocketServer webSocketServer `toml:"webSocketServer"`
	SQLLog          sqlLog          `toml:"sql_log"`
	// SSL             ssl             `toml:"ssl"`
	SSL struct {
		CertPath string
		KeyPath  string
	} `toml:"ssl"`
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

type logType struct {
	FileName         string
	NbFilesToArchive int
}

type server struct {
	Port         int
	AllowOrigins []string
}

type webSocketServer struct {
	Port int
}

type ssl struct {
	certPath string
	keyPath  string
}

type sqlLog struct {
	Level            uint
	Limit            float64
	DisplayOverLimit bool
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
