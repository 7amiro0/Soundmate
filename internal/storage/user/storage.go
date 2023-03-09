package storage

import (
	"context"
	"social_network/internal/storage"

	"github.com/jackc/pgx/v4"
)

const (
	add = "insert into users (name, email, password) values ($1, $2, $3) returning id;"
	getByName = "select * from users where name=$1"
	getByEmail = "select * from users where email=$1"
	checkUsers = "select * from users where email=$1 and password=$2"
	// list = ""
	// update = "update users set name=$1, email=$2, password=$3 where id=$4"
	// delete = "delete from users where id=$1"
)

type UserStorage struct {
	conn *pgx.Conn
	config storage.StorageConfig
	ctx context.Context
}

func NewStorageUsers(ctx context.Context, config storage.StorageConfig) *UserStorage {
	return &UserStorage{
		config: config,
		ctx: ctx,
	}
}

func (s *UserStorage) Connect() (err error) {
	s.conn, err = pgx.Connect(s.ctx, s.config.GetLink())
	return err
}

func (s *UserStorage) Disconnect() error {
	return s.conn.Close(s.ctx)
}

func (s *UserStorage) Add(user *storage.User) error {
	rows := s.conn.QueryRow(s.ctx, add, user.Name, user.Email, user.Password)
	if err := rows.Scan(&user.ID); err != nil {
		return err
	}

	return nil
}

func (s *UserStorage) GetByName(name string) storage.Users {
	var users storage.Users
	rows := s.conn.QueryRow(s.ctx, getByName, name)
	rows.Scan(users)

	return users
}

func (s *UserStorage) GetByEmail(email []byte) storage.Users {
	var users storage.Users
	rows := s.conn.QueryRow(s.ctx, getByEmail, email)
	rows.Scan(users)

	return users
}

func (s *UserStorage) CheckUsers(email, password []byte) storage.Users {
	var users storage.Users
	rows := s.conn.QueryRow(s.ctx, checkUsers, email, password)
	rows.Scan(users)

	return users
}
