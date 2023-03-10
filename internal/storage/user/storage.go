package storage

import (
	"context"
	"social_network/internal/storage"

	"github.com/jackc/pgx/v4"
)

const (
	add        = "insert into users (name, email, password) values ($1, $2, $3) returning id;"
	getByName  = "select * from users where name=$1;"
	getByEmail = "select * from users where email=$1;"
	checkUsers = "select * from users where email=$1 and password=$2;"
)

type UserStorage struct {
	conn   *pgx.Conn
	config storage.StorageConfig
	ctx    context.Context
}

func NewStorageUsers(ctx context.Context, config storage.StorageConfig) *UserStorage {
	return &UserStorage{
		config: config,
		ctx:    ctx,
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

func list(rows pgx.Rows) (storage.Users, error) {
	var users []storage.User
	for rows.Next() {
		var user storage.User
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
		)

		if err != nil {
			return storage.Users{}, err
		}

		users = append(users, user)
	}

	return storage.Users{Users: users}, nil
}

func (s *UserStorage) GetByName(name string) (storage.Users, error) {
	rows, err := s.conn.Query(s.ctx, getByName, name)
	if err != nil {
		return storage.Users{}, err
	}

	return list(rows)
}

func (s *UserStorage) GetByEmail(email []byte) (storage.Users, error) {
	rows, err := s.conn.Query(s.ctx, getByEmail, email)
	if err != nil {
		return storage.Users{}, err
	}

	return list(rows)
}

func (s *UserStorage) CheckUsers(email, password []byte) (storage.Users, error) {
	rows, err := s.conn.Query(s.ctx, checkUsers, email, password)
	if err != nil {
		return storage.Users{}, err
	}

	return list(rows)
}
