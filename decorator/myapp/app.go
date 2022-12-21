package myapp

import (
	"fmt"
	"net/http"
)

func indexHandelr(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandelr)
	return mux
}
