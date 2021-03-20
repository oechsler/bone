package main

import (
	"github.com/labstack/echo/v4"
	"github.com/oechsler/bone/posts"
	"go.uber.org/dig"
)

func main() {
	var err error

	server := echo.New()
	container := dig.New()

	err = container.Provide(func() *echo.Echo {
		return server
	})
	if err != nil {
		server.Logger.Fatal(err)
	}

	err = container.Provide(func() echo.Logger {
		return server.Logger
	})
	if err != nil {
		server.Logger.Fatal(err)
	}

	// Install modules
	err = posts.UseModule(container);
	if err != nil {
		server.Logger.Fatal(err)
	}

	err = server.Start(":1323")
	if err != nil {
		server.Logger.Fatal(err)
	}
}
