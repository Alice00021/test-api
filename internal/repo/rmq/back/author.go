package back

import (
	"context"
	"github.com/Alice00021/test_api/config"
	"github.com/Alice00021/test_api/internal/entity/back"
	"github.com/Alice00021/test_common/pkg/rabbitmq/rmq_rpc/client"
)

type AuthorRMQ struct {
	*client.Client
	Receivers config.RMQReceivers
}

func NewAuthorRMQ(client *client.Client, receivers config.RMQReceivers) *AuthorRMQ {
	return &AuthorRMQ{client, receivers}
}
func (m *AuthorRMQ) CreateAuthor(ctx context.Context, inp back.CreateAuthorInput) (*back.Author, error) {
	var resp *back.Author
	err := m.RemoteCall(ctx, m.Receivers.BackService, "v1.createAuthor", inp, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *AuthorRMQ) UpdateAuthor(ctx context.Context, inp back.UpdateAuthorInput) error {
	err := m.RemoteCall(ctx, m.Receivers.BackService, "v1.updateAuthor", inp, nil)
	if err != nil {
		return err
	}

	return nil
}

func (m *AuthorRMQ) GetAuthor(ctx context.Context, id int64) (*back.Author, error) {
	var resp *back.Author

	req := make(map[string]int64)
	req["id"] = id

	err := m.RemoteCall(ctx, m.Receivers.BackService, "v1.getAuthor", req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *AuthorRMQ) GetAuthors(ctx context.Context) ([]*back.Author, error) {
	var resp []*back.Author

	err := m.RemoteCall(ctx, m.Receivers.BackService, "v1.getAuthors", nil, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *AuthorRMQ) DeleteAuthor(ctx context.Context, id int64) error {
	req := make(map[string]int64)
	req["id"] = id

	err := m.RemoteCall(ctx, m.Receivers.BackService, "v1.deleteAuthor", req, nil)
	if err != nil {
		return err
	}

	return nil
}
