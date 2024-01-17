package main

import (
	"github.com/joho/godotenv"
	"github.com/sanstzu/pocket-api/src/api"
	"github.com/sanstzu/pocket-api/src/consts"
	watch "github.com/sanstzu/pocket-api/src/services/watch"
)

func main() {
	var err error

	err = godotenv.Load()
	if err != nil {
		panic(err)
	}

	consts.Init()

	err = watch.StartClient()
	if err != nil {
		panic(err)
	}

	api.Start()
}
