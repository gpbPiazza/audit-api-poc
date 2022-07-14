package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gpbPiazza/src/interfaces"
	"log"
)

type HttpServer struct {
	Middlewares interfaces.MiddlewaresContainer
	FiberApp    *fiber.App
}

const ENV_VAR = 3001

func NewHttpServer(middlewares interfaces.MiddlewaresContainer) *HttpServer {
	return &HttpServer{
		Middlewares: middlewares,
		FiberApp:    fiber.New(),
	}
}

func (hs *HttpServer) Start() {
	hs.setUpMiddlewares(hs.Middlewares)

	if err := hs.FiberApp.Listen(fmt.Sprintf(":%v", ENV_VAR)); err != nil {
		log.Fatalf("server Err=%v", err.Error())
	}
}

func (hs *HttpServer) setUpMiddlewares(middlewares interfaces.MiddlewaresContainer) {
	hs.FiberApp.Use(middlewares.Cors)
	hs.FiberApp.Use(middlewares.ContentType)
}
