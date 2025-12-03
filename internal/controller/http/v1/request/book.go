package request

import "github.com/Alice00021/test_api/internal/entity/back"

type CreateBookRequest struct {
	Title    string `json:"title" validate:"required"`
	AuthorId int64  `json:"authorId" validate:"required"`
}

func (req *CreateBookRequest) ToEntity() back.CreateBookInput {
	return back.CreateBookInput{
		Title:    req.Title,
		AuthorId: req.AuthorId,
	}
}

type UpdateBookRequest struct {
	CreateBookRequest
}

func (req *UpdateBookRequest) ToEntity() back.UpdateBookInput {
	return back.UpdateBookInput{
		Title:    req.Title,
		AuthorId: req.AuthorId,
	}
}
