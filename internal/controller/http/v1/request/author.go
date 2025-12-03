package request

import "github.com/Alice00021/test_api/internal/entity/back"

type CreateAuthorRequest struct {
	Name   string `json:"name" validate:"required"`
	Gender bool   `json:"gender" validate:"required"`
}

func (req *CreateAuthorRequest) ToEntity() back.CreateAuthorInput {
	return back.CreateAuthorInput{
		Name:   req.Name,
		Gender: req.Gender,
	}
}

type UpdateAuthorRequest struct {
	Name string `json:"name" validate:"required"`
}

func (req *UpdateAuthorRequest) ToEntity() back.UpdateAuthorInput {
	return back.UpdateAuthorInput{
		Name: req.Name,
	}
}
