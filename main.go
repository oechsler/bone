package main

import (
	"github.com/labstack/echo/v4"
	"github.com/oechsler/bone/middleware"
	"github.com/oechsler/bone/posts"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

func main() {
	var err error

	server := echo.New()
	container := dig.New()

	// Configure middleware
	logger := logrus.New()
	server.Logger = middleware.NewLogrusAdapter(logger)

	// Add server to the dependency injection
	err = container.Provide(func() *echo.Echo {
		return server
	})
	if err != nil {
		server.Logger.Fatal(err)
	}

	// Add logger to the dependency injection
	err = container.Provide(func() echo.Logger {
		return server.Logger
	})
	if err != nil {
		server.Logger.Fatal(err)
	}

	// Add modules to the dependency injection
	err = posts.UseModule(container)
	if err != nil {
		server.Logger.Fatal(err)
	}

	// Start the server
	err = server.Start(":1323")
	if err != nil {
		server.Logger.Fatal(err)
	}
}
