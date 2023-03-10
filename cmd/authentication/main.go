package main

import (
	"context"
	"log"
	"net"
	"os/signal"
	"social_network/internal/servers/registration"
	"social_network/internal/storage/user"
	"syscall"
)

func main() {
	config := NewConfig()

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	userStorage := storage.NewStorageUsers(ctx, config.db)

	server := registration.New(
		userStorage,
		net.JoinHostPort(config.server.host, config.server.port),
	)

	err := server.Connect()
	if err != nil {
		log.Fatal("Cant start server:", err)
	}

	<-ctx.Done()

	err = server.Disconnect()
	if err != nil {
		log.Fatal("Cant stop server:", err)
	}
}
