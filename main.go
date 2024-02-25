package main

import (
	"fmt"

	"literank.com/rest-books/adaptor"
	"literank.com/rest-books/application"
	"literank.com/rest-books/infrastructure/config"
)

const configFileName = "config.yml"

func main() {
	// Read the config
	c, err := config.Parse(configFileName)
	if err != nil {
		panic(err)
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
