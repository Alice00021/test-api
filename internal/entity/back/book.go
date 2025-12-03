package back

import "github.com/Alice00021/test_api/internal/entity"

type Book struct {
	entity.Entity
	Title    string
	AuthorId int64
	Author   Author
}

type CreateBookInput struct {
	Title    string `json:"name"`
	AuthorId int64  `json:"author_id"`
}

type UpdateBookInput struct {
	ID       int64  `json:"id"`
	Title    string `json:"name"`
	AuthorId int64  `json:"author_id"`
}
