package storage

import (
	"fmt"
	"os"
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
	Name     string
	Email    []byte
	Password []byte
	ID       uint
}

type Users struct {
	Users []User
}
