package usecase

import (
	"context"
	"github.com/Alice00021/test_api/internal/entity/back"
)

type (
	Auth interface {
		Register(context.Context, back.CreateUserInput) (*back.User, error)
		Login(context.Context, back.AuthenticateInput) (*back.TokenPair, error)
		VerifyEmail(context.Context, back.VerifyEmail) error
	}

	Author interface {
		CreateAuthor(context.Context, back.CreateAuthorInput) (*back.Author, error)
		UpdateAuthor(context.Context, back.UpdateAuthorInput) error
		GetAuthors(context.Context) ([]*back.Author, error)
		GetAuthor(context.Context, int64) (*back.Author, error)
		DeleteAuthor(context.Context, int64) error
	}

	Book interface {
		CreateBook(context.Context, back.CreateBookInput) (*back.Book, error)
		UpdateBook(context.Context, back.UpdateBookInput) error
		GetBook(context.Context, int64) (*back.Book, error)
		GetBooks(context.Context) ([]*back.Book, error)
		DeleteBook(context.Context, int64) error
	}

	Command interface {
		UpdateCommands(context.Context) error
		GetCommands(context.Context) ([]*back.Command, error)
	}

	Operation interface {
		CreateOperation(context.Context, back.CreateOperationInput) (*back.Operation, error)
		UpdateOperation(context.Context, back.UpdateOperationInput) error
		GetOperations(context.Context) ([]*back.Operation, error)
		GetOperation(context.Context, int64) (*back.Operation, error)
		DeleteOperation(context.Context, int64) error
	}
)
