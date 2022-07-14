package handlers

import (
	"github.com/gpbPiazza/src/interfaces"
)

type Handler interface {
	Routes(middlewares interfaces.MiddlewaresContainer)
}

type HandlerContainer struct {
}

func NewHandlerContainer() HandlerContainer {
	return HandlerContainer{}
}
