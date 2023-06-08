package user

import (
	"html/template"
	"io"
	"log"
	"mime/multipart"
	"social_network/internal/storage"
	"strings"

	"github.com/gofiber/fiber/v2"
)

const (
	userHTML    = "./templates/user.html"
	musicHTML   = "./templates/add_music.html"
	settingHTML = "./templates/setting_user.html"
)

func (s *Server) user(c *fiber.Ctx) error {
	email := c.Cookies("email")
	if email == "" {
		return c.Redirect("/auth/login")
	}

	user, err := s.userStorage.GetByEmail([]byte(email))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	musics, err := s.musicStorage.GetByAuthor(user[0].ID)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	tmpl := template.Must(template.ParseFiles(userHTML))
	c.Response().Header.Set("Content-Type", "text/html")
	err = tmpl.Execute(c.Response().BodyWriter(), fiber.Map{
		"UserName":    user[0].Name,
		"Musics":      musics,
		"Description": user[0].Description,
	})
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return err
}

func (s *Server) settingUser(c *fiber.Ctx) error {
	email := c.Cookies("email")
	if email == "" {
		return c.Redirect("/auth/login")
	}

	user, err := s.userStorage.GetByEmail([]byte(email))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if c.Method() == fiber.MethodGet {
		return c.Render(settingHTML, fiber.Map{
			"UserName":    user[0].Name,
			"Description": user[0].Description,
		})
	} else if c.Method() == fiber.MethodPost {
		name := strings.TrimSpace(c.FormValue("name"))
		if len(name) < 4 {
			c.SendStatus(fiber.StatusUnprocessableEntity)
			return c.Render(settingHTML, nil)
		}

		user[0].Name = name
		user[0].Description = c.FormValue("description")

		err = s.userStorage.Update(user[0])
		if err != nil {
			c.SendStatus(fiber.StatusInternalServerError)
			return err
		}

		return c.Render(settingHTML, fiber.Map{
			"UserName":    user[0].Name,
			"Description": user[0].Description,
		})
	} else {
		return nil
	}
}

func readFile(file multipart.File) ([]byte, error) {
	return io.ReadAll(file)
}

func (s *Server) addMusic(c *fiber.Ctx) error {
	email := c.Cookies("email")
	if email == "" {
		return c.Redirect("/auth/login")
	}

	user, err := s.userStorage.GetByEmail([]byte(email))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if c.Method() == fiber.MethodGet {
		return c.Render(musicHTML, nil)
	} else if c.Method() == fiber.MethodPost {
		title := c.FormValue("title")
		
		musicFile, err := c.FormFile("music_file")
		if err != nil {
			log.Println("get music file", err)
		}

		photoFile, err := c.FormFile("photo")
		if err != nil {
			log.Println("get photo file", err)
		}

		file, err := musicFile.Open()
		if err != nil {
			log.Println("open music file", err)
		}
		defer file.Close()
		
		byteMusicFile, err := readFile(file)
		if err != nil {
			log.Println("read music file", err)
		}

		file, err = photoFile.Open()
		if err != nil {
			log.Println("open photo file", err)
		}
		defer file.Close()
		
		bytePhotoFile, err := readFile(file)
		if err != nil {
			log.Println("read photo file", err)
		}

		err = s.musicStorage.Add(storage.Music{
			Photo:        bytePhotoFile,
			File:         byteMusicFile,
			Author:       user[0].Name,
			Title:        title,
			AuthorID:     user[0].ID,
		})

		if err != nil {
			log.Println("add file", err)
		}

		return c.Redirect("/user")
	} else {
		return nil
	}
}
