package interfaces

import "social_network/internal/storage"

type ConnecterDisconnecter interface {
	Connect() error
	Disconnect() error
}

type StorageUser interface {
	ConnecterDisconnecter
	Add(user *storage.User) error
	GetByName(name string) (storage.Users, error)
	GetByEmail(email []byte) (storage.Users, error)
	CheckUsers(email, password []byte) (storage.Users, error)
}
