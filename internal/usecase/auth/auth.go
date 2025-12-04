package auth

import (
	"context"
	"github.com/Alice00021/test_api/internal/entity/back"
	"github.com/Alice00021/test_api/internal/repo/rmq"
)

type UseCase struct {
	rmq rmq.AuthRMQ
}

func New(rmq rmq.AuthRMQ) *UseCase {
	return &UseCase{rmq}
}

func (u *UseCase) Register(ctx context.Context, inp back.CreateUserInput) (*back.User, error) {
	return u.rmq.Register(ctx, inp)
}

func (u *UseCase) Login(ctx context.Context, inp back.AuthenticateInput) (*back.TokenPair, error) {
	return u.rmq.Login(ctx, inp)
}

func (u *UseCase) VerifyEmail(ctx context.Context, inp back.VerifyEmail) error {
	return u.rmq.VerifyEmail(ctx, inp)
}

func (u *UseCase) Validation(ctx context.Context, accessToken string) (*back.UserInfoToken, error) {
	return u.rmq.Validation(ctx, accessToken)
}
