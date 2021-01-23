package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/flowdev/example-micro/trolley-cmd/service"
	"github.com/flowdev/example-micro/trolley-cmd/web"
)

func main() {
	fmt.Println("main()")
	router := web.NewRouter(service.New())
	log.Fatal(http.ListenAndServe("8080", router))
}
