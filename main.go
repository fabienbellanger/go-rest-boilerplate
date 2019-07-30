package main

import (
	"time"

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

	time.Sleep(10 * time.Second)
}
