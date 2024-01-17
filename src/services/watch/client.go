package runner

import (
	pb "github.com/sanstzu/pocket-api/internal/watch"
	"github.com/sanstzu/pocket-api/src/consts"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client *pb.WatchClient

func getClient() *pb.WatchClient {
	return client
}

func StartClient() error {
	var opts []grpc.DialOption

	// TODO: Implement TLS security for gRPC
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	config := consts.GetConsts().JudgeRPCConfig

	conn, err := grpc.Dial(config.ServerAddress, opts...)

	if err != nil {
		return err
	}

	tmp := pb.NewWatchClient(conn)

	client = &tmp

	return nil
}
