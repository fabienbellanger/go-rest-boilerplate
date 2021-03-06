package lib

import (
	"os"
	"reflect"
	"strconv"
	"strings"
)

// CheckError manages errors
func CheckError(err error, exitCode int) {
	if err != nil {
		if exitCode != 0 {
			CustomLog("Error(" + strconv.Itoa(exitCode) + "): " + err.Error())

			os.Exit(exitCode)
		} else {
			CustomLog(err.Error())
		}
	}
}

// Ucfirst makes a string's first character uppercase
func Ucfirst(s string) string {
	// Tableau de caractères Unicode pour gérér les caractères accentués
	sToUnicode := []rune(s)

	return strings.ToUpper(string(sToUnicode[0])) + string(sToUnicode[1:])
}

// InArray search an element in an array
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
