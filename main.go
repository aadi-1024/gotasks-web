package main

import (
	"github.com/aadi-1024/gotasks-web/api"
	"log"
	"net/http"
)

func main() {
	err := api.AddRoutes(http.DefaultServeMux)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(http.ListenAndServe(":8080", nil))
}
