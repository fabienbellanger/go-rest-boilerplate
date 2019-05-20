package websockets

import (
	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"net/http"
	"strconv"
)

// ServerStart starts websockets server
func ServerStart(port int) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ClientConnection(w, r)
	})

	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	lib.CheckError(err, -1)
}
