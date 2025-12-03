package di

import (
	"github.com/Alice00021/test_api/config"
	backRMQ "github.com/Alice00021/test_api/internal/repo/rmq/back"
	"github.com/Alice00021/test_api/internal/usecase"
	"github.com/Alice00021/test_api/internal/usecase/auth"
	"github.com/Alice00021/test_api/internal/usecase/author"
	"github.com/Alice00021/test_api/internal/usecase/book"
	"github.com/Alice00021/test_api/internal/usecase/command"
	"github.com/Alice00021/test_api/internal/usecase/operation"
	"github.com/Alice00021/test_common/pkg/rabbitmq/rmq_rpc/client"
)

type UseCase struct {
	Auth      usecase.Auth
	Author    usecase.Author
	Book      usecase.Book
	Command   usecase.Command
	Operation usecase.Operation
}

func NewUseCase(rmqClient *client.Client, conf *config.Config) *UseCase {
	return &UseCase{
		Auth:      auth.New(backRMQ.NewAuthRMQ(rmqClient, conf.RMQReceivers)),
		Author:    author.New(backRMQ.NewAuthorRMQ(rmqClient, conf.RMQReceivers)),
		Book:      book.New(backRMQ.NewBookRMQ(rmqClient, conf.RMQReceivers)),
		Command:   command.New(backRMQ.NewCommandRMQ(rmqClient, conf.RMQReceivers)),
		Operation: operation.New(backRMQ.NewOperationRMQ(rmqClient, conf.RMQReceivers)),
	}
}
