package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/flowdev/example-micro/externalacts-qry/service"
	"github.com/flowdev/example-micro/externalacts-qry/web"
)

func main() {
	fmt.Println("main()")
	router := web.NewRouter(service.New())
	log.Fatal(http.ListenAndServe("8080", router))
}
