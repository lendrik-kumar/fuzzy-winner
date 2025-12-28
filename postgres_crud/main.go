package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lendrik-kumar/postgres-crud/router"
)

func main() {

	r := router.Router()
	fmt.Println("starting server on 6000")

	log.Fatal(http.ListenAndServe(":6000", r))

}

