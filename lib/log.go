package lib

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/viper"
)

// DefaultEchoLogWriter displays errors logs on the good writer
var DefaultEchoLogWriter io.Writer

// DefaultSqlLogWriter displays SQL logs on the good writer
var DefaultSqlLogWriter io.Writer

var (
	redColor   = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	resetColor = string([]byte{27, 91, 48, 109})
)

// DisplaySuccessMessage displays success message to output
func DisplaySuccessMessage(msg string) {
	if len(msg) > 0 {
		green := color.New(color.FgGreen).SprintFunc()
		fmt.Printf(" %s  %s\n", green("✔"), msg)
	}
}

// CustomLog displays logs
func CustomLog(v ...interface{}) {
	if DefaultEchoLogWriter == nil {
		DefaultEchoLogWriter = os.Stdout
	}
	log.SetOutput(DefaultEchoLogWriter)

	// Remove logs timestamp
	log.SetFlags(0)

	log.Printf("ERR  | %s | %+v\n", time.Now().Format("2006-01-02 15:04:05"), v)
}

// SQLLog displays SQL log in gin.DefaultWriter
func SQLLog(latency time.Duration, query string, args ...interface{}) {
	if DefaultSqlLogWriter == nil {
		DefaultSqlLogWriter = os.Stdout
	}
	log.SetOutput(DefaultSqlLogWriter)

	// Remove logs timestamp
	log.SetFlags(0)

	// Traitement de la requête
	// ------------------------
	query = strings.Join(strings.Fields(query), " ")

	// Couleur (en environnement de développement)
	// -------------------------------------------
	var latencyColor string

	if viper.GetString("environment") == "production" {
		latencyColor = ""
		resetColor = ""
	} else {
		latencyColor = resetColor

		if latency.Seconds() >= viper.GetFloat64("log.sql.limit") {
			latencyColor = redColor
		}
	}

	// Type de requête
	// ---------------
	queryArray := strings.Fields(query)
	requestType := ""

	if len(queryArray) > 0 {
		requestType = strings.ToUpper(queryArray[0])
		requestType = requestType[0:3]
	}

	// Affichage des logs
	// ------------------
	if viper.GetInt("log.sql.level") == 1 {
		// Time only
		// ---------
		log.Printf("SQL  | %s | %4s |%s %v %s\t|\n",
			time.Now().Format("2006-01-02 15:04:05"),
			requestType,
			latencyColor, latency, resetColor)
	} else if viper.GetInt("log.sql.level") == 2 {
		// Time and query
		// --------------
		log.Printf("SQL  | %s | %4s |%s %v %s\t| %s |\n",
			time.Now().Format("2006-01-02 15:04:05"),
			requestType,
			latencyColor, latency, resetColor,
			query)
	} else if viper.GetInt("log.sql.level") == 3 {
		// Time, query and arguments
		// -------------------------
		log.Printf("SQL  | %s | %4s |%s %v %s\t| %s | %v\n",
			time.Now().Format("2006-01-02 15:04:05"),
			requestType,
			latencyColor, latency, resetColor,
			query,
			args)
	}
}
