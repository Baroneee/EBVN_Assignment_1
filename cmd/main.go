package main

import "github.com/Baroneee/EBVN_Assignment_1/internal/api"

func main() {
	cfg, err := api.NewConfig()
	if err != nil {
		panic(err)
	}

	app := api.NewEngine(cfg)
	err = app.Start()
	if err != nil {
		panic(err)
	}
}