package main

import (
	"WebDemo/Server"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test", r.URL.Path[1:])
}

func main() {
	Server.StartCustomHandlerServer()
}
