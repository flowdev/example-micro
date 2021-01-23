package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/flowdev/example-micro/hedwig-cmd/service"
	"github.com/flowdev/example-micro/hedwig-cmd/web"
)

func main() {
	fmt.Println("main()")
	router := web.NewRouter(service.New())
	log.Fatal(http.ListenAndServe("8080", router))
}