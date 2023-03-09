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
	// Photo     string
	// Email    string
}

func (s *Server) home(c *fiber.Ctx) error {
	email := c.Cookies("email")
	// if email == "" {
		// return c.Redirect("/auth/login")
	// }

	// user := s.storage.GetByEmail([]byte(email)).Users[0]

	return c.Render(homeHTML, info{
		UserName: "some " + string(email),
	})
}

func (s *Server) user(c *fiber.Ctx) error {
	email := c.Cookies("email")
	if email == "" {
		return c.Redirect("/auth/login")
	}

	user := s.storage.GetByEmail([]byte(email)).Users[0]

	return c.Render(userHTML, info{
		UserName: user.Name,
	})
}