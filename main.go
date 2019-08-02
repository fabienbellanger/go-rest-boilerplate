package main

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/fabienbellanger/go-rest-boilerplate/issues"
)

func main() {
	// Lancement de Cobra
	// commands.Execute()
	go issues.InitData()
	go issues.InitData()
	go issues.InitData()
	go issues.InitData()
	go issues.InitData()

	// time.Sleep(60 * time.Second)

	http.ListenAndServe("localhost:8082", nil)
}
