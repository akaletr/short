package app

import (
	"cmd/short/main.go/internal/storage"
	"fmt"
	"log/slog"
)

type App struct {
	log     *slog.Logger
	storage storage.InterfaceStorage
}

func (a *App) InitApp() {

}

func (a *App) StartApp() {
	const op = "app.StartApp"

	url, err := a.storage.SaveURL("https://www.youtube.com/watch?v=rCJvW2xgnk0")
	if err != nil {
		fmt.Println(url, err)
	}

}

func (a *App) StopApp() {

}

func New(log *slog.Logger, storage storage.InterfaceStorage) InterfaceApp {
	return &App{
		log:     log,
		storage: storage,
	}
}
