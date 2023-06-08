package authentication

import (
	"social_network/internal/storage"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"

	"crypto/sha256"
	"encoding/hex"
)

const (
	// HTML file
	loginHTML        = "./templates/login.html"
	registrationHTML = "./templates/registration.html"
	notFoundHTML     = "./templates/notFound.html"

	// Error messages for html
	errEmailOrPassword = "Not correct email or password"
	errNameRepeat      = "This name is already using"
	errNameNotCorrect  = "This name isn't correct"
	errEmail           = "This email is already using"
)

func hide(data string) ([]byte, error) {
	h := sha256.New()
	_, err := h.Write([]byte(data))
	if err != nil {
		return nil, err
	}

	dst := make([]byte, hex.EncodedLen(len(h.Sum(nil))))

	hex.Encode(dst, h.Sum(nil))
	return dst, nil
}

func (s *Server) registration(c *fiber.Ctx) error {
	if c.Method() == fiber.MethodGet {
		return c.Render(registrationHTML, nil)
	} else if c.Method() == fiber.MethodPost {
		name := strings.TrimSpace(c.FormValue("name"))
		if len(name) < 4 {
			return c.Render(registrationHTML, fiber.Map{
				"Name": errNameNotCorrect,
			})
		}

		users, err := s.storage.GetByName(name)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		if len(users) != 0 {
			return c.Render(registrationHTML, fiber.Map{
				"Name": errNameRepeat,
			})
		}

		email, err := hide(c.FormValue("email"))
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		users, err = s.storage.GetByEmail(email)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		if len(users) != 0 {
			return c.Render(registrationHTML, fiber.Map{
				"Email": errEmail,
			})
		}

		password, err := hide(c.FormValue("password"))
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		user := storage.User{
			Email:    email,
			Password: password,
			Name:     name,
		}

		err = s.storage.Add(user)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		c.Cookie(&fiber.Cookie{
			Name:     "email",
			Value:    string(email),
			Expires:  time.Now().Add(time.Hour * 24),
			HTTPOnly: true,
		})

		return c.Redirect("/home")
	}

	return nil
}

func (s *Server) login(c *fiber.Ctx) error {
	if c.Method() == fiber.MethodGet {
		return c.Render(loginHTML, nil)
	} else if c.Method() == fiber.MethodPost {
		email, err := hide(c.FormValue("email"))
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		password, err := hide(c.FormValue("password"))
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		user, err := s.storage.CheckUsers(email, password)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		if len(user) == 0 {
			return c.Render(loginHTML, fiber.Map{
				"EmailOrPassword": errEmailOrPassword,
			})
		}

		c.Cookie(&fiber.Cookie{
			Name:     "email",
			Value:    string(email),
			Expires:  time.Now().Add(time.Hour * 24),
			HTTPOnly: true,
		})

		return c.Redirect("/home")
	}

	return nil
}

func (s *Server) NotFound(c *fiber.Ctx) error {
	return c.Render(notFoundHTML, nil)
}
