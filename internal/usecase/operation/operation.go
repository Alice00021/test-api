package operation

import (
	"context"
	"github.com/Alice00021/test_api/internal/entity/back"
	"github.com/Alice00021/test_api/internal/repo/rmq"
)

type UseCase struct {
	rmq rmq.OperationRMQ
}

func New(rmq rmq.OperationRMQ) *UseCase {
	return &UseCase{rmq}
}

func (u *UseCase) GetOperations(ctx context.Context) ([]*back.Operation, error) {
	return u.rmq.GetOperations(ctx)
}

func (u *UseCase) GetOperation(ctx context.Context, id int64) (*back.Operation, error) {
	return u.rmq.GetOperation(ctx, id)
}

func (u *UseCase) CreateOperation(ctx context.Context, inp back.CreateOperationInput) (*back.Operation, error) {
	return u.rmq.CreateOperation(ctx, inp)
}

func (u *UseCase) UpdateOperation(ctx context.Context, inp back.UpdateOperationInput) error {
	return u.rmq.UpdateOperation(ctx, inp)
}

func (u *UseCase) DeleteOperation(ctx context.Context, id int64) error {
	return u.rmq.DeleteOperation(ctx, id)
}
