package v1

import "SSEHandler/ssehandler"

type handler struct {
	sseHandler *ssehandler.SSEHandler
}

func New(sseHandler *ssehandler.SSEHandler) *handler {
	return &handler{
		sseHandler: sseHandler,
	}
}
