package back

import (
	"context"
	"github.com/Alice00021/test_api/config"
	"github.com/Alice00021/test_api/internal/entity/back"
	"github.com/Alice00021/test_common/pkg/rabbitmq/rmq_rpc/client"
)

type BookRMQ struct {
	*client.Client
	Receivers config.RMQReceivers
}

func NewBookRMQ(client *client.Client, receivers config.RMQReceivers) *BookRMQ {
	return &BookRMQ{client, receivers}
}
func (m *BookRMQ) CreateBook(ctx context.Context, inp back.CreateBookInput) (*back.Book, error) {
	var resp *back.Book
	err := m.RemoteCall(ctx, m.Receivers.BackService, "v1.createBook", inp, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *BookRMQ) UpdateBook(ctx context.Context, inp back.UpdateBookInput) error {
	err := m.RemoteCall(ctx, m.Receivers.BackService, "v1.updateBook", inp, nil)
	if err != nil {
		return err
	}

	return nil
}

func (m *BookRMQ) GetBook(ctx context.Context, id int64) (*back.Book, error) {
	var resp *back.Book

	req := make(map[string]int64)
	req["id"] = id

	err := m.RemoteCall(ctx, m.Receivers.BackService, "v1.getBook", req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *BookRMQ) GetBooks(ctx context.Context) ([]*back.Book, error) {
	var resp []*back.Book

	err := m.RemoteCall(ctx, m.Receivers.BackService, "v1.getBooks", nil, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *BookRMQ) DeleteBook(ctx context.Context, id int64) error {
	req := make(map[string]int64)
	req["id"] = id

	err := m.RemoteCall(ctx, m.Receivers.BackService, "v1.deleteBook", req, nil)
	if err != nil {
		return err
	}

	return nil
}
