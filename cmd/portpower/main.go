package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"portpower/internal/console"
	"portpower/internal/ns"
	"portpower/internal/store"
)

func main() {
	dir := os.Getenv("PORTPOWER_DATA_DIR")
	if dir == "" {
		dir = "data"
	}
	st, err := store.Open(dir)
	if err != nil {
		log.Fatalf("open store: %v", err)
	}
	harbor := ns.NewHarbor("pp-north", "port-north")
	server := console.NewServer(harbor, st)
	addr := os.Getenv("PORTPOWER_ADDR")
	if addr == "" {
		addr = ":8080"
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	serverHTTP := &http.Server{Addr: addr, Handler: server.Handler()}
	go func() {
		<-ctx.Done()
		_ = serverHTTP.Shutdown(context.Background())
	}()
	log.Printf("portpower listening on %s", addr)
	if err := serverHTTP.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("serve: %v", err)
	}
}
