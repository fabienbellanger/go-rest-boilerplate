package lib

import (
	"strings"

	"github.com/spf13/viper"
)

// InitConfig : Lecture du fichier de configuration
func InitConfig(file string) {
	viper.SetConfigFile(file)
	err := viper.ReadInConfig()
	CheckError(err, 1)

	// On converti le / du timezone de la base de donnée en %2F
	// --------------------------------------------------------
	viper.Set(
		"database.timezone",
		strings.Replace(viper.GetString("database.timezone"), "/", "%2F", -1))
}

// IsDatabaseConfigCorrect : La configuration de la base de données est-elle correcte ?
func IsDatabaseConfigCorrect() bool {
	return viper.GetString("database.driver") != "" &&
		viper.GetString("database.name") != "" &&
		viper.GetString("database.host") != ""
}

// IsServerConfigCorrect checks if server config is correct
func IsServerConfigCorrect() bool {
	return len(viper.GetStringSlice("server.allowOrigins")) > 0
}
