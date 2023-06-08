package home

import (
	"github.com/gofiber/fiber/v2"
)

const (
	// HTML file
	homeHTML = "./templates/home.html"
)


func (s *Server) home(c *fiber.Ctx) error {
	email := c.Cookies("email")
	if email == "" {
		return c.Redirect("/auth/login")
	}

	user, err := s.storage.GetByEmail([]byte(email))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Render(homeHTML, fiber.Map{
		"UserName": "some " + user[0].Name,
	})
}
