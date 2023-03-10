package main

import (
	"os"
	"social_network/internal/storage"
)

type Config struct {
	db     storage.StorageConfig
	server ServerConfig
}

type ServerConfig struct {
	host string
	port string
}

func (s *ServerConfig) Set() {
	s.host = os.Getenv("HOST")
	s.port = os.Getenv("PORT_HOME")
}

func (c *Config) Set() {
	c.db.Set()
	c.server.Set()
}

func NewConfig() *Config {
	config := &Config{}
	config.Set()

	return config
}
