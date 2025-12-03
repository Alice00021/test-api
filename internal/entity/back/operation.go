package back

import "github.com/Alice00021/test_api/internal/entity"

type OperationCommand struct {
	ID          int64
	OperationID int64
	Command
	Address Address
}

type Operation struct {
	entity.Entity
	Name        string
	Description string
	AverageTime int64
	Commands    []*OperationCommand
}

type UpdateOperationInput struct {
	ID          int64           `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Commands    []*CommandInput `json:"commands"`
}

type CommandInput struct {
	ID         *int64
	SystemName string
	Address    Address
}

type CreateOperationInput struct {
	Name        string
	Description string
	AverageTime int64
	Commands    []*CommandInput
}
