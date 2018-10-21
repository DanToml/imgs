package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var (
	httpAddr string
)

func init() {
	httpAddr = os.Getenv("HTTP_ADDR")
	if httpAddr == "" {
		httpAddr = "0.0.0.0:8080"
	}
}

func main() {
	r := chi.NewRouter()

	r.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Recoverer,
		middleware.GetHead,
		middleware.Timeout(2*time.Second),
		middleware.Logger,
	)

	r.Route("/api", func(r chi.Router) {
		r.Use(
			middleware.AllowContentType("application/json"),
			middleware.SetHeader("Content-Type", "application/json"),
		)

		r.Get("/ping", func(rw http.ResponseWriter, req *http.Request) {
			rw.Write([]byte("{}"))
		})
	})

	srv := &http.Server{
		Addr: httpAddr,

		WriteTimeout:      15 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
		IdleTimeout:       60 * time.Second,

		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}
