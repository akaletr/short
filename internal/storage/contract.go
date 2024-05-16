package storage

import "errors"

type InterfaceStorage interface {
	Run()

	SaveURL(url string) (string, error)
	GetURL(alias string) (string, error)
	DeleteURL(alias string) error
}

var (
	ErrURLNotFound = errors.New("url not found")
	ErrURLExists   = errors.New("url exists")
)
