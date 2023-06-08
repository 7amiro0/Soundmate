package home

import (
	"social_network/internal/interfaces"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	server  *fiber.App
	storage interfaces.StorageUser
	addr    string
}

func New(storage interfaces.StorageUser, addr string) *Server {
	serverFiber := fiber.New(fiber.Config{})

	return &Server{
		server:  serverFiber,
		storage: storage,
		addr:    addr,
	}
}

func (s *Server) setRouter() {
	s.server.Get("/home", s.home)
}

func (s *Server) Connect() error {
	s.setRouter()

	return s.server.Listen(s.addr)
}

func (s *Server) Disconnect() error {
	return s.server.Shutdown()
}
