package consts

// Creatres config singleton
type Consts struct {
	JudgeRPCConfig JudgeRPC
}

var config *Consts

func Init() {
	config = &Consts{
		JudgeRPCConfig: initializeJudgeRpc(),
	}
}

func GetConsts() *Consts {
	return config
}
