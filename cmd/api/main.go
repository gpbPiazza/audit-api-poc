package main

import (
	"github.com/gpbPiazza/src/http"
	"github.com/gpbPiazza/src/middlewares"
)

func main() {
	middlewaresContainer := middlewares.NewMiddlewaresContainer()

	server := http.NewHttpServer(middlewaresContainer)

	server.Start()
}
