package main

import (
	"errors"
	"net/http"
	_ "net/http/pprof"

	"github.com/fabienbellanger/go-rest-boilerplate/issues"

	"github.com/fabienbellanger/go-rest-boilerplate/database"
	"github.com/fabienbellanger/go-rest-boilerplate/lib"
)

func main() {
	lib.InitConfig("config.toml")
	if !lib.IsDatabaseConfigCorrect() {
		err := errors.New("no or missing database information in settings file")
		panic(err)
		// lib.CheckError(err, 1)
	}

	database.Open()
	defer database.DB.Close()

	// Lancement de Cobra
	// commands.Execute()
	go issues.InitData()
	// go issues.InitData()
	// go issues.InitData()
	// go issues.InitData()
	// go issues.InitData()

	http.ListenAndServe("localhost:8082", nil)
}
