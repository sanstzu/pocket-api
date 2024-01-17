package consts

import (
	"os"
)

type JudgeRPC struct {
	ServerAddress string
	JudgeTimeout  int
}

func initializeJudgeRpc() JudgeRPC {
	var config JudgeRPC

	config.ServerAddress = os.Getenv("JUDGE_RPC_SERVER_ADDRESS")
	config.JudgeTimeout = 10000

	return config
}
