package rmq

import (
	"context"
	"github.com/Alice00021/test_api/internal/entity/back"
)

type (
	AuthRMQ interface {
		Register(context.Context, back.CreateUserInput) (*back.User, error)
		Login(context.Context, back.AuthenticateInput) (*back.TokenPair, error)
		VerifyEmail(context.Context, back.VerifyEmail) error
		Validation(context.Context, string) (*back.UserInfoToken, error)
	}

	AuthorRMQ interface {
		CreateAuthor(context.Context, back.CreateAuthorInput) (*back.Author, error)
		UpdateAuthor(context.Context, back.UpdateAuthorInput) error
		GetAuthor(context.Context, int64) (*back.Author, error)
		GetAuthors(context.Context) ([]*back.Author, error)
		DeleteAuthor(context.Context, int64) error
	}

	BookRMQ interface {
		CreateBook(context.Context, back.CreateBookInput) (*back.Book, error)
		UpdateBook(context.Context, back.UpdateBookInput) error
		GetBook(context.Context, int64) (*back.Book, error)
		GetBooks(context.Context) ([]*back.Book, error)
		DeleteBook(context.Context, int64) error
	}

	CommandRMQ interface {
		UpdateCommands(context.Context) error
		GetCommands(context.Context) ([]*back.Command, error)
	}

	OperationRMQ interface {
		CreateOperation(context.Context, back.CreateOperationInput) (*back.Operation, error)
		UpdateOperation(context.Context, back.UpdateOperationInput) error
		GetOperation(context.Context, int64) (*back.Operation, error)
		GetOperations(context.Context) ([]*back.Operation, error)
		DeleteOperation(context.Context, int64) error
	}
)
