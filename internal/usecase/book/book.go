package book

import (
	"context"
	"github.com/Alice00021/test_api/internal/entity/back"
	"github.com/Alice00021/test_api/internal/repo/rmq"
)

type UseCase struct {
	rmq rmq.BookRMQ
}

func New(rmq rmq.BookRMQ) *UseCase {
	return &UseCase{rmq}
}

func (u *UseCase) GetBooks(ctx context.Context) ([]*back.Book, error) {
	return u.rmq.GetBooks(ctx)
}

func (u *UseCase) GetBook(ctx context.Context, id int64) (*back.Book, error) {
	return u.rmq.GetBook(ctx, id)
}

func (u *UseCase) CreateBook(ctx context.Context, inp back.CreateBookInput) (*back.Book, error) {
	return u.rmq.CreateBook(ctx, inp)
}

func (u *UseCase) UpdateBook(ctx context.Context, inp back.UpdateBookInput) error {
	return u.rmq.UpdateBook(ctx, inp)
}

func (u *UseCase) DeleteBook(ctx context.Context, id int64) error {
	return u.rmq.DeleteBook(ctx, id)
}
