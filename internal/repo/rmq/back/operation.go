package back

import (
	"context"
	"github.com/Alice00021/test_api/config"
	"github.com/Alice00021/test_api/internal/entity/back"
	"github.com/Alice00021/test_common/pkg/rabbitmq/rmq_rpc/client"
)

type OperationRMQ struct {
	*client.Client
	Receivers config.RMQReceivers
}

func NewOperationRMQ(client *client.Client, receivers config.RMQReceivers) *OperationRMQ {
	return &OperationRMQ{client, receivers}
}
func (m *OperationRMQ) CreateOperation(ctx context.Context, inp back.CreateOperationInput) (*back.Operation, error) {
	var resp *back.Operation
	err := m.RemoteCall(ctx, m.Receivers.BackService, "v1.createOperation", inp, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *OperationRMQ) UpdateOperation(ctx context.Context, inp back.UpdateOperationInput) error {
	err := m.RemoteCall(ctx, m.Receivers.BackService, "v1.updateOperation", inp, nil)
	if err != nil {
		return err
	}

	return nil
}

func (m *OperationRMQ) GetOperation(ctx context.Context, id int64) (*back.Operation, error) {
	var resp *back.Operation

	req := make(map[string]int64)
	req["id"] = id

	err := m.RemoteCall(ctx, m.Receivers.BackService, "v1.getOperation", req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *OperationRMQ) GetOperations(ctx context.Context) ([]*back.Operation, error) {
	var resp []*back.Operation

	err := m.RemoteCall(ctx, m.Receivers.BackService, "v1.getOperations", nil, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *OperationRMQ) DeleteOperation(ctx context.Context, id int64) error {
	req := make(map[string]int64)
	req["id"] = id

	err := m.RemoteCall(ctx, m.Receivers.BackService, "v1.deleteOperation", req, nil)
	if err != nil {
		return err
	}

	return nil
}
