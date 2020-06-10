package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
)

func main(){
	index := os.Getenv("POD_NAME")
	http.HandleFunc("/get", func(writer http.ResponseWriter, request *http.Request) {
		i := rand.Intn(5)
		if i <=3 {
			writer.WriteHeader(http.StatusBadGateway)
		}
		fmt.Fprintf(writer,"I am %s, random i = %d\n",index,i)
	})
	http.ListenAndServe(":38080",nil)
}
