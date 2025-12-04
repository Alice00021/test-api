package v1

import (
	"github.com/Alice00021/test_api/internal/di"
	"github.com/Alice00021/test_common/pkg/logger"
	"github.com/Alice00021/test_common/pkg/rabbitmq/rmq_rpc/server"
)

func NewRouter(routes map[string]server.CallHandler, uc *di.UseCase, l logger.Interface) {

}
