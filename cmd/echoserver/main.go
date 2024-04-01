package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"echoserver/internal/echoserver"
)

func main() {

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	if err := echoserver.Run(ctx); err != nil {
		log.Fatalf("error can't start echo server: %v", err)
	}

}
