package pg

import (
	"cmd/short/main.go/internal/domain/models"
	"cmd/short/main.go/internal/lib"
	"cmd/short/main.go/internal/storage"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	conn *gorm.DB
}

func (s *Storage) Run() {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) SaveURL(url string) (string, error) {
	u := models.URL{
		Alias: "https://www.youtube.com",
		URL:   url,
	}

	s.conn.Create(&u)
	return u.Alias, nil
}

func (s *Storage) GetURL(alias string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) DeleteURL(alias string) error {
	//TODO implement me
	panic("implement me")
}

func New(db string) (storage.InterfaceStorage, error) {
	const op = "storage.pg.New"
	conn, err := gorm.Open(postgres.Open(db), &gorm.Config{})
	if err != nil {
		return nil, lib.Err(op, err)
	}

	return &Storage{conn: conn}, nil
}
