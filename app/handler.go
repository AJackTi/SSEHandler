package app

import (
	v1 "SSEHandler/controller/http/v1"
	"SSEHandler/httpserver"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"log"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	handler := gin.New()
	pprof.Register(handler, "dev/pprof")
	handler.Use(cors.Default())
	v1.NewRouter(handler)

	httpServer := httpserver.New(handler, httpserver.Port("8081"))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Println("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		log.Printf("app - Run - httpServer.Notify %v", err)
	}

	// Shutdown
	err := httpServer.Shutdown()
	if err != nil {
		log.Printf("app - Run - httpServer.Shutdown %v", err)
	}
}
