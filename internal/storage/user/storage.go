package user

import (
	"context"
	"social_network/internal/storage"

	"github.com/jackc/pgx/v4"
)

const (
	add        = "insert into users (name, email, password) values ($1, $2, $3);"
	update     = "update users set name=$1, description=$2 where id=$3;"
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

func (s *UserStorage) Add(user storage.User) error {
	_, err := s.conn.Exec(s.ctx, add, user.Name, user.Email, user.Password)
	return err
}

func list(rows pgx.Rows) ([]storage.User, error) {
	var users []storage.User
	for rows.Next() {
		var (
			user        storage.User
			description *string
		)
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&description,
			&user.Email,
			&user.Password,
			&user.Photo,
		)

		if err != nil {
			return nil, err
		}

		if description != nil {
			user.Description = *description
		}

		users = append(users, user)
	}

	return users, nil
}

func (s *UserStorage) Update(user storage.User) error {
	_, err := s.conn.Exec(s.ctx, update, user.Name, user.Description, user.ID)
	return err
}

func (s *UserStorage) GetByName(name string) ([]storage.User, error) {
	rows, err := s.conn.Query(s.ctx, getByName, name)
	if err != nil {
		return nil, err
	}

	return list(rows)
}

func (s *UserStorage) GetByEmail(email []byte) ([]storage.User, error) {
	rows, err := s.conn.Query(s.ctx, getByEmail, email)
	if err != nil {
		return nil, err
	}

	return list(rows)
}

func (s *UserStorage) CheckUsers(email, password []byte) ([]storage.User, error) {
	rows, err := s.conn.Query(s.ctx, checkUsers, email, password)
	if err != nil {
		return nil, err
	}

	return list(rows)
}
