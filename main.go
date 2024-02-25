package main

import (
	"fmt"

	"literank.com/rest-books/adaptor"
	"literank.com/rest-books/application"
	"literank.com/rest-books/infrastructure/config"
)

func main() {
	c := &config.Config{
		App: config.ApplicationConfig{
			Port: 8080,
		},
		DB: config.DBConfig{
			FileName: "test.db",
		},
	}
	// Prepare dependencies
	wireHelper, err := application.NewWireHelper(c)
	if err != nil {
		panic(err)
	}

	// Build main router
	r, err := adaptor.MakeRouter(wireHelper)
	if err != nil {
		panic(err)
	}
	// Run the server on the specified port
	if err := r.Run(fmt.Sprintf(":%d", c.App.Port)); err != nil {
		panic(err)
	}
}
