package music

import (
	"context"
	"encoding/base64"
	"social_network/internal/storage"

	"github.com/jackc/pgx/v4"
)

const (
	add = "insert into musics (authorID, title, author, music, photo) values ($1, $2, $3, $4, $5);"
	getByAuthor = "select * from musics where authorID=$1;"
	getByID = "select * from musics where id=$1;"
	getByNameAndAuhtor = "select * from musics where title=$1 and authorID=$2"
)

type MusicStorage struct {
	conn   *pgx.Conn
	config storage.StorageConfig
	ctx    context.Context
}

func NewMusicStorage(ctx context.Context, config storage.StorageConfig) *MusicStorage {
	return &MusicStorage{
		config: config,
		ctx:    ctx,
	}
}

func (s *MusicStorage) Connect() (err error) {
	s.conn, err = pgx.Connect(s.ctx, s.config.GetLink())
	return err
}

func (s *MusicStorage) Disconnect() error {
	return s.conn.Close(s.ctx)
}

func (s *MusicStorage) Add(music storage.Music) error {
	_, err := s.conn.Exec(s.ctx, add, music.AuthorID, music.Title, music.Author, music.File, music.Photo)
	return err
}

func list(rows pgx.Rows) ([]storage.Music, error) {
	var musics []storage.Music
	for rows.Next() {
		var (
			music storage.Music
		)

		err := rows.Scan(
			&music.ID,
			&music.AuthorID,
			&music.Title,
			&music.Author,
			&music.File,
			&music.Photo,
		)
		if err != nil {
			return nil, err
		}

		music.EncodedFile = base64.StdEncoding.EncodeToString(music.File)
		music.EncodedPhoto = base64.StdEncoding.EncodeToString(music.Photo)

		musics = append(musics, music)
	}

	return musics, nil
}

func (s *MusicStorage) GetByID(id uint) ([]storage.Music, error) {
	rows, err := s.conn.Query(s.ctx, getByID, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return list(rows)
}

func (s *MusicStorage) GetByAuthor(authorID uint) ([]storage.Music, error) {
	rows, err := s.conn.Query(s.ctx, getByAuthor, authorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return list(rows)
}