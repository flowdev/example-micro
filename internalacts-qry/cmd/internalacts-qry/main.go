package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/flowdev/example-micro/internalacts-qry/service"
	"github.com/flowdev/example-micro/internalacts-qry/web"
)

func main() {
	fmt.Println("main()")
	router := web.NewRouter(service.New())
	log.Fatal(http.ListenAndServe("8080", router))
}
