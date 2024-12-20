package server

import (
	"fmt"
	"net/http"
)

// Run server to communicate with DB
func RunServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "WriteOutDB v0.0.0")
	})	

	http.ListenAndServe(":63030", nil)
}
