package main

import (
	_ "net/http/pprof"

	"github.com/fabienbellanger/go-rest-boilerplate/commands"
)

func main() {
	// lib.InitConfig("config.toml")
	// if !lib.IsDatabaseConfigCorrect() {
	// 	err := errors.New("no or missing database information in settings file")
	// 	panic(err)
	// 	// lib.CheckError(err, 1)
	// }

	// database.Open()
	// defer database.DB.Close()

	// Lancement de Cobra
	commands.Execute()
	// go issues.InitData()
	// go issues.InitData()
	// go issues.InitData()
	// go issues.InitData()
	// go issues.InitData()

	// http.ListenAndServe("localhost:8082", nil)
}
