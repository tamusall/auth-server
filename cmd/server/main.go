package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/tamusall/auth-server/internal/transport/http"
)

type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting up things")

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8081", handler.Router); err != nil {
		fmt.Println("Failed to configure setup server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("hello there")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting server")
		fmt.Println(err)
	}
}
