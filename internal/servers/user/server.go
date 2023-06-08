package user

import (
	"social_network/internal/interfaces"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	server  *fiber.App
	userStorage interfaces.StorageUser
	musicStorage interfaces.StorageMusic
	addr    string
}

func New(musicStorage interfaces.StorageMusic, userStorage interfaces.StorageUser, addr string) *Server {
	return &Server{
		server:  fiber.New(fiber.Config{}),
		userStorage: userStorage,
		musicStorage: musicStorage,
		addr:    addr,
	}
}

func (s *Server) setRouter() {
	s.server.Get("/user", s.user)
	
	s.server.Get("/user/setting", s.settingUser)
	s.server.Post("/user/setting", s.settingUser)

	s.server.Get("/user/music", s.addMusic)
	s.server.Post("/user/music", s.addMusic)
}

func (s *Server) Connect() error {
	s.setRouter()

	return s.server.Listen(s.addr)
}

func (s *Server) Disconnect() error {
	return s.server.Shutdown()
}
