package back

import (
	"context"
	"github.com/Alice00021/test_api/config"
	"github.com/Alice00021/test_api/internal/entity/back"
	"github.com/Alice00021/test_common/pkg/rabbitmq/rmq_rpc/client"
)

type CommandRMQ struct {
	*client.Client
	Receivers config.RMQReceivers
}

func NewCommandRMQ(client *client.Client, receivers config.RMQReceivers) *CommandRMQ {
	return &CommandRMQ{client, receivers}
}

func (m *CommandRMQ) UpdateCommands(ctx context.Context) error {
	err := m.RemoteCall(ctx, m.Receivers.BackService, "v1.updateCommands", nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (m *CommandRMQ) GetCommands(ctx context.Context) ([]*back.Command, error) {
	var resp []*back.Command

	err := m.RemoteCall(ctx, m.Receivers.BackService, "v1.getCommands", nil, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
