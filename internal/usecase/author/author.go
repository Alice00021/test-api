package author

import (
	"context"
	"github.com/Alice00021/test_api/internal/entity/back"
	"github.com/Alice00021/test_api/internal/repo/rmq"
)

type UseCase struct {
	rmq rmq.AuthorRMQ
}

func New(rmq rmq.AuthorRMQ) *UseCase {
	return &UseCase{rmq}
}

func (u *UseCase) GetAuthors(ctx context.Context) ([]*back.Author, error) {
	return u.rmq.GetAuthors(ctx)
}

func (u *UseCase) GetAuthor(ctx context.Context, id int64) (*back.Author, error) {
	return u.rmq.GetAuthor(ctx, id)
}

func (u *UseCase) CreateAuthor(ctx context.Context, inp back.CreateAuthorInput) (*back.Author, error) {
	return u.rmq.CreateAuthor(ctx, inp)
}

func (u *UseCase) UpdateAuthor(ctx context.Context, inp back.UpdateAuthorInput) error {
	return u.rmq.UpdateAuthor(ctx, inp)
}

func (u *UseCase) DeleteAuthor(ctx context.Context, id int64) error {
	return u.rmq.DeleteAuthor(ctx, id)
}
