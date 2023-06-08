package main

import (
	"context"
	"log"
	"net"
	"os/signal"
	"social_network/internal/servers/authentication"
	"social_network/internal/storage/user"
	"syscall"
)

func main() {
	config := NewConfig()

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	userStorage := user.NewStorageUsers(ctx, config.db)
	err := userStorage.Connect()
	if err != nil {
		log.Fatalln("Cant connect to storage:", err)
	}
	defer func() {
		err = userStorage.Disconnect()
		if err != nil {
			log.Fatalln("Cant disconnect storage:", err)
		}
	}()

	server := authentication.New(
		userStorage,
		net.JoinHostPort(config.server.host, config.server.port),
	)

	err = server.Connect()
	if err != nil {
		log.Fatalln("Cant start server:", err)
	}

	<-ctx.Done()

	err = server.Disconnect()
	if err != nil {
		log.Fatalln("Cant stop server:", err)
	}
}
