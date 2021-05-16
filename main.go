package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ibrahim-akrab/products_api/handlers"
)

func main() {
	l := log.New(os.Stdout, "products_api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, os.Kill)

	sig := <-sigChan
	l.Printf("received signal %s, starting graceful shutdown\n", sig)
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	shutdownErr := s.Shutdown(ctx)
	l.Println(shutdownErr)
}
