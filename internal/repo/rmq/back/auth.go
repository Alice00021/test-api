package back

import (
	"context"
	"github.com/Alice00021/test_api/config"
	"github.com/Alice00021/test_api/internal/entity/back"
	"github.com/Alice00021/test_common/pkg/rabbitmq/rmq_rpc/client"
)

type AuthRMQ struct {
	*client.Client
	Receivers config.RMQReceivers
}

func NewAuthRMQ(client *client.Client, receivers config.RMQReceivers) *AuthRMQ {
	return &AuthRMQ{client, receivers}
}

func (m *AuthRMQ) Register(ctx context.Context, inp back.CreateUserInput) (*back.User, error) {
	var resp *back.User

	err := m.RemoteCall(ctx, m.Receivers.TxService, "v1.register", inp, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *AuthRMQ) Login(ctx context.Context, inp back.AuthenticateInput) (*back.TokenPair, error) {
	var resp *back.TokenPair

	err := m.RemoteCall(ctx, m.Receivers.TxService, "v1.login", inp, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *AuthRMQ) VerifyEmail(ctx context.Context, inp back.VerifyEmail) error {

	err := m.RemoteCall(ctx, m.Receivers.TxService, "v1.verifyEmail", inp, nil)
	if err != nil {
		return err
	}
	return nil
}
