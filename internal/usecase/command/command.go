package command

import (
	"context"
	"github.com/Alice00021/test_api/internal/entity/back"
	"github.com/Alice00021/test_api/internal/repo/rmq"
)

type UseCase struct {
	rmq rmq.CommandRMQ
}

func New(rmq rmq.CommandRMQ) *UseCase {
	return &UseCase{rmq}
}

func (u *UseCase) UpdateCommands(ctx context.Context) error {
	return u.rmq.UpdateCommands(ctx)
}

func (u *UseCase) GetCommands(ctx context.Context) ([]*back.Command, error) {
	return u.rmq.GetCommands(ctx)
}
