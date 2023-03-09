package app

import "social_network/internal/storage"

type ConnecterDisconnecter interface {
	Connect() error
	Disconnect() error
}

type StorageUser interface {
	ConnecterDisconnecter
	Add(user *storage.User) error
	GetByName(name string) storage.Users
	GetByEmail(email []byte) storage.Users
	CheckUsers(email, password []byte) storage.Users
}
