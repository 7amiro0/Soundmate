package interfaces

import "social_network/internal/storage"

type StorageUser interface {
	Add(user storage.User) error
	Update(user storage.User) error
	GetByName(name string) ([]storage.User, error)
	GetByEmail(email []byte) ([]storage.User, error)
	CheckUsers(email, password []byte) ([]storage.User, error)
}

type StorageMusic interface {
	Add(music storage.Music) error
	GetByAuthor(id uint) ([]storage.Music, error)
	GetByID(id uint) ([]storage.Music, error)
	// Delete()
	// Update()
}
