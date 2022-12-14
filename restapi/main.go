package main

import (
	"go-toy-web/restapi/myapp"
	"net/http"
)

func main() {
	http.ListenAndServe(":3200", myapp.NewHandler())
}
