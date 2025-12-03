package back

import "github.com/Alice00021/test_api/internal/entity"

type Author struct {
	entity.Entity
	Name   string
	Gender bool
}

type CreateAuthorInput struct {
	Name   string `json:"name"`
	Gender bool   `json:"gender"`
}

type UpdateAuthorInput struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
