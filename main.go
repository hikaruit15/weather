package main

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	apiKeyChecker := NewApiKeyChecker()
	h := NewHandler(&mockStore{})
	router := gin.New()

	// add middleware
	router.Use(cors.Default())
	router.Use(gin.Logger())
	router.Use(apiKeyChecker)

	// add handler
	router.GET("/nearest_city", h.GetNearestCity)
	router.GET("/countries", h.GetCountries)
	router.GET("/states", h.GetStates)
	router.GET("/cities", h.GetCities)

	// create http server and start it
	srv := &http.Server{Addr: ":8088", Handler: router}
	StartServerWithGracefulShutdown(srv)
}

func StartServerWithGracefulShutdown(srv *http.Server) {
	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
