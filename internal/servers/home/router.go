package registration

import (
	"github.com/gofiber/fiber/v2"
)

const (
	// HTML file
	homeHTML = "./templates/home.html"
	userHTML = "./templates/user.html"
)

type info struct {
	UserName string
}

func (s *Server) home(c *fiber.Ctx) error {
	email := c.Cookies("email")
	if email == "" {
		return c.Redirect("/auth/login")
	}

	user, err := s.storage.GetByEmail([]byte(email))
	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
		return err
	}

	return c.Render(homeHTML, info{
		UserName: "some " + user.Users[0].Name,
	})
}

func (s *Server) user(c *fiber.Ctx) error {
	email := c.Cookies("email")
	if email == "" {
		return c.Redirect("/auth/login")
	}

	user, err := s.storage.GetByEmail([]byte(email))
	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
		return err
	}

	return c.Render(homeHTML, info{
		UserName: "some " + user.Users[0].Name,
	})
}
