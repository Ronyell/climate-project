package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()
	fmt.Println(fmt.Sprintf("Run server in port %d!", config.Port))
	r := router.Generate()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", rune(config.Port)), r))
}
