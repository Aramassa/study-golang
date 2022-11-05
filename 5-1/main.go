package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, "Hello world!")

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)

	http.ListenAndServe(":8080", handlers.CompressHandler(r))
}
