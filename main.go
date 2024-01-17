package main

import (
	"github.com/joho/godotenv"
	"github.com/sanstzu/pocket-api/src/api"
	"github.com/sanstzu/pocket-api/src/consts"
	judgeRpc "github.com/sanstzu/pocket-api/src/services/judge-rpc"
)

func main() {
	var err error

	err = godotenv.Load()
	if err != nil {
		panic(err)
	}

	consts.Init()

	err = judgeRpc.StartClient()
	if err != nil {
		panic(err)
	}

	api.Start()
}
