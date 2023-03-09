package registration

import (
	"log"
	"social_network/internal/storage"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const (
	// HTML file
	loginHTML = "./templates/login.html"
	registrationHTML = "./templates/registration.html"

	// Error message for html
	errEmailOrPassword = "Not correct email or password"
	errNameRepeat = "This name is already using"
	errNameNotCorrect = "This name isn't correct"
	errEmail = "This email is already using"
)

type errMessages struct {
	EmailOrPassword string
	Name string
	Email string
}

func hide(data string) ([]byte, error) {
	return bcrypt.GenerateFromPassword(
		[]byte(data),
		bcrypt.DefaultCost,
	)
}

func (s *Server) registration(c *fiber.Ctx) error {
	log.Println("sign up")
	if c.Method() == "GET" {
		return c.Render(registrationHTML, nil)
	} else if c.Method() == "POST" {
		name := strings.TrimSpace(c.FormValue("name"))
		if len(name) < 4 /*|| !match*/ {
			c.SendStatus(fiber.StatusUnprocessableEntity)
			return c.Render(loginHTML, errMessages{
				Name: errNameNotCorrect,
			})
		}

		users := s.storage.GetByName(name)
		if len(users.Users) != 0 {
			c.SendStatus(fiber.StatusUnprocessableEntity)
			return c.Render(loginHTML, errMessages{
				Name: errNameRepeat,
			})
		}

		email, err := hide(c.FormValue("email"))
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		users = s.storage.GetByEmail(email)
		if len(users.Users) != 0 {
			c.SendStatus(fiber.StatusUnprocessableEntity)
			return c.Render(loginHTML, errMessages{
				Email: errEmail,
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

		err = s.storage.Add(&user)
		if err != nil {
			c.SendStatus(fiber.StatusInternalServerError)
			return err
		}

		c.Cookie(&fiber.Cookie{
			Name:     "email",
			Value:    string(email),
			Expires:  time.Now().Add(time.Hour * 24),
			HTTPOnly: true,
		})

		return c.Redirect("/home")
	} else {
		return c.SendStatus(fiber.StatusNotFound)
	}
}

func (s *Server) login(c *fiber.Ctx) error {
	if c.Method() == "GET" {
		return c.Render(loginHTML, nil)
	} else if c.Method() == "POST" {
		email, err := hide(c.FormValue("email"))
		if err != nil {
			c.SendStatus(fiber.StatusInternalServerError)
			return err
		}

		password, err := hide(c.FormValue("password"))
		if err != nil {
			c.SendStatus(fiber.StatusInternalServerError)
			return err
		}

		users := s.storage.CheckUsers(email, password)
		if len(users.Users) == 0 {
			return c.Render(loginHTML, errMessages{
				EmailOrPassword: errEmailOrPassword,
			})
		}



		c.Cookie(&fiber.Cookie{
			Name:     "email",
			Value:    string(email),
			Expires:  time.Now().Add(time.Hour * 24),
			HTTPOnly: true,
		})

		return c.Redirect("/home")
	} else {
		return c.SendStatus(fiber.StatusNotFound)
	}
}
