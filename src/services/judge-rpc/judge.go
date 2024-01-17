package judge

import (
	"context"
	"log"
	"time"

	pb "github.com/sanstzu/pocket-api/internal/judge"
	"github.com/sanstzu/pocket-api/src/api/models"
	"github.com/sanstzu/pocket-api/src/consts"
)

func Judge(req *models.JudgeRequest) (*models.JudgeResponse, error) {
	client := getClient()
	consts := consts.GetConsts().JudgeRPCConfig

	rpcReq := &pb.JudgeRequest{
		Code:     req.Code,
		Language: req.Language,
		Input:    req.Stdin,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(consts.JudgeTimeout)*time.Millisecond)

	defer cancel()

	res, err := (*client).Judge(ctx, rpcReq)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	resModel := &models.JudgeResponse{
		Verdict: res.Verdict,
		Stdout:  res.Stdout,
		Stderr:  res.Stderr,
		Time:    res.Time,
		Memory:  res.Memory,
	}

	return resModel, err
}
