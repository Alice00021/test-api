package request

import "github.com/Alice00021/test_api/internal/entity/back"

type CreateOperationRequest struct {
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Commands    []*back.CommandInput `json:"commands"`
}

func (req *CreateOperationRequest) ToEntity() back.CreateOperationInput {
	return back.CreateOperationInput{
		Name:        req.Name,
		Description: req.Description,
		Commands:    req.Commands,
	}
}

type UpdateOperationRequest struct {
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Commands    []*back.CommandInput `json:"commands"`
}

func (req *UpdateOperationRequest) ToEntity() back.UpdateOperationInput {
	return back.UpdateOperationInput{
		Name:        req.Name,
		Description: req.Description,
		Commands:    req.Commands,
	}
}
