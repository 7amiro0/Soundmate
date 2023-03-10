package registration

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
	s.server.Get("/auth/login", s.login)
	s.server.Post("/auth/login", s.login)

	s.server.Get("/auth/registration", s.registration)
	s.server.Post("/auth/registration", s.registration)
}

func (s *Server) Connect() error {
	err := s.storage.Connect()
	if err != nil {
		return err
	}

	s.setRouter()

	return s.server.Listen(s.addr)
}

func (s *Server) Disconnect() error {
	err := s.storage.Disconnect()
	if err != nil {
		return err
	}

	return s.server.Shutdown()
}
