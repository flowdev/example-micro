package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/flowdev/example-micro/galton-qry/service"
	"github.com/flowdev/example-micro/galton-qry/web"
)

func main() {
	fmt.Println("main()")
	router := web.NewRouter(service.New())
	log.Fatal(http.ListenAndServe("8080", router))
}
