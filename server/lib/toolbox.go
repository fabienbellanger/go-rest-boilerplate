package lib

import (
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// CheckError : Gestion des erreurs
func CheckError(err error, exitCode int) {
	if err != nil {
		if exitCode != 0 {
			log.Println("Error (" + strconv.Itoa(exitCode) + "): " + err.Error() + "\n\n")

			os.Exit(exitCode)
		} else {
			log.Println(err.Error() + "\n")
		}
	}
}

// Ucfirst : Met la première lettre d'une chaîne de caractères en majuscule
func Ucfirst(s string) string {
	sToUnicode := []rune(s) // Tableau de caractères Unicode pour gérér les caractères accentués

	return strings.ToUpper(string(sToUnicode[0])) + string(sToUnicode[1:])
}

// InArray : Recherche dans un tableau
func InArray(value interface{}, array interface{}) (found bool, index int) {
	index = -1
	found = false

	switch reflect.Indirect(reflect.ValueOf(array)).Kind() {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(array)
		sLen := s.Len()

		for i := 0; i < sLen; i++ {
			if reflect.DeepEqual(value, s.Index(i).Interface()) {
				index = i
				found = true

				break
			}
		}
	}

	return
}
