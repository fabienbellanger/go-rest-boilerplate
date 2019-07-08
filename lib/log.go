package lib

import (
	"io"
	"log"
	"strings"
	"time"
)

// DefaultEchoLogWriter displays logs on the good writer
var DefaultEchoLogWriter io.Writer

var (
	redColor   = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	resetColor = string([]byte{27, 91, 48, 109})
)

// CustomLog displays logs
func CustomLog(v ...interface{}) {
	// On redirige les logs vers le default writer de Gin
	// log.SetOutput(gin.DefaultWriter)
	log.SetOutput(DefaultEchoLogWriter)

	log.Printf("ERR  | %s | %+v\n", time.Now().Format(time.RFC3339), v)

	// mutex := new(sync.Mutex)
	// mutex.Lock()
	// go func(v ...interface{}) {
	// 	log.Printf("ERR  | %s | %+v\n", time.Now().Format(time.RFC3339), v)
	// 	mutex.Unlock()
	// }(v)
}

// SQLLog displays SQL log in gin.DefaultWriter
func SQLLog(latency time.Duration, query string, args ...interface{}) {
	// On redirige les logs vers le default writer de Gin
	// log.SetOutput(gin.DefaultWriter)
	log.SetOutput(DefaultEchoLogWriter)

	// Remove logs timestamp
	log.SetFlags(0)

	// Traitement de la requête
	// ------------------------
	query = strings.Join(strings.Fields(query), " ")

	// Couleur (en environnement de développement)
	// -------------------------------------------
	var latencyColor string

	if Config.Environment == "production" {
		latencyColor = ""
		resetColor = ""
	} else {
		latencyColor = resetColor

		if Config.Environment == "development" && latency.Seconds() >= Config.SQLLog.Limit {
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
	if Config.SQLLog.Level == 1 {
		// Time only
		// ---------
		log.Printf("SQL  | %s | %4s |%s %13v %s|\n",
			// time.Now().Format("2006/01/02 - 15:04:05"),
			time.Now().Format(time.RFC3339),
			requestType,
			latencyColor, latency, resetColor)
	} else if Config.SQLLog.Level == 2 {
		// Time and query
		// --------------
		log.Printf("SQL  | %s | %4s |%s %13v %s| %s\n",
			// time.Now().Format("2006/01/02 - 15:04:05"),
			time.Now().Format(time.RFC3339),
			requestType,
			latencyColor, latency, resetColor,
			query)
	} else if Config.SQLLog.Level == 3 {
		// Time, query and arguments
		// -------------------------
		log.Printf("SQL  | %s | %4s |%s %13v %s| %s | %v\n",
			// time.Now().Format("2006/01/02 - 15:04:05"),
			time.Now().Format(time.RFC3339),
			requestType,
			latencyColor, latency, resetColor,
			query,
			args)
	}
}
