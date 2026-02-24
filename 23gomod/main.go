package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", ServeHome)

	http.ListenAndServe("localhost:4000", r)

}

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Welcome </h1>"))
}
