package lib

import (
	"apiticSellers/server/lib"
	"os"
)

const LOG_FILENAME = "gin.log"

// ExecuteLogsRotation launches logs rotation
//
// Description:
// Every day at 00:00
func ExecuteLogsRotation() {
	_, err := os.OpenFile(LOG_FILENAME, os.O_RDONLY, 0755)

	if err != nil {
		lib.CheckError(err, -1)
	}

	// Décalage des fichiers du type LOG_FILENAME.x
	// --------------------------------------------

	// Déplacement du fichier de log
	// -----------------------------
	err = os.Rename(LOG_FILENAME, LOG_FILENAME+".1")

	if err != nil {
		lib.CheckError(err, -2)
	}

	// Création du nouveau fichier LOG_FILENAME
	// ----------------------------------------
	_, err = os.Create(LOG_FILENAME)

	if err != nil {
		// Le fichier de log n'existe pas
		// ------------------------------
		lib.CheckError(err, -3)
	}
}
