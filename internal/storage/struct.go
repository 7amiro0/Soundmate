package storage

import (
	"fmt"
	"os"
	"strconv"
)

type StorageConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func (s StorageConfig) GetLink() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		s.Host,
		s.Port,
		s.User,
		s.Password,
		s.DBName,
	)
}

func (d *StorageConfig) Set() {
	d.Host = os.Getenv("DATABASE_HOST")
	d.Port = os.Getenv("DATABASE_PORT")
	d.User = os.Getenv("POSTGRES_USER")
	d.DBName = os.Getenv("POSTGRES_DB")
	d.Password = os.Getenv("POSTGRES_PASSWORD")
}

type User struct {
	Photo        []byte
	Email        []byte
	Password     []byte
	EncodedPhoto string
	Name         string
	Description  string
	ID           uint
}

type Users struct {
	Users []User
}

func (m Music) GetTitle() string {
	return m.Title
}

func (m Music) GetStringID() string {
	return strconv.Itoa(int(m.ID))
}

type Music struct {
	Photo        []byte
	File         []byte
	EncodedFile  string
	EncodedPhoto string
	Author       string
	Title        string
	AuthorID     uint
	ID           uint
}
