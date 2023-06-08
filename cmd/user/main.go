package main

import (
	"context"
	"log"
	"net"
	"os/signal"
	"social_network/internal/storage/music"
	"social_network/internal/storage/user"
	userServer "social_network/internal/servers/user"
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
		log.Fatalln("Cant connect to user storage:", err)
	}
	defer func() {
		err = userStorage.Disconnect()
		if err != nil {
			log.Fatalln("Cant disconnect user storage:", err)
		}
	}()
	
	musicStorage := music.NewMusicStorage(ctx, config.db)
	err = musicStorage.Connect()
	if err != nil {
		log.Fatalln("Cant connect to music storage:", err)
	}
	defer func() {
		err = musicStorage.Disconnect()
		if err != nil {
			log.Fatalln("Cant disconnect music storage:", err)
		}
	}()

	server := userServer.New(
		musicStorage,
		userStorage,
		net.JoinHostPort(config.server.host, config.server.port),
	)

	err = server.Connect()
	if err != nil {
		log.Fatal("Cant start server:", err)
	}

	<-ctx.Done()

	err = server.Disconnect()
	if err != nil {
		log.Fatal("Cant stop server:", err)
	}
}