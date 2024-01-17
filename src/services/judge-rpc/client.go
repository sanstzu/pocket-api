package judge

import (
	pb "github.com/sanstzu/pocket-api/internal/judge"
	"github.com/sanstzu/pocket-api/src/consts"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client *pb.JudgeClient

func getClient() *pb.JudgeClient {
	return client
}

func StartClient() error {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	config := consts.GetConsts().JudgeRPCConfig

	conn, err := grpc.Dial(config.ServerAddress, opts...)

	if err != nil {
		return err
	}

	tmp := pb.NewJudgeClient(conn)

	client = &tmp

	return nil
}
