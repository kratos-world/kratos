package template

var (
	Service = `package service

import "github.com/google/wire"

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(New{{.BCName}}Service)`

	CustomService = `package service

import (
	"context"
	"errors"
	"fmt"

	v1 "{{.ModPath}}/api/{{.SCName}}"
	"{{.ModPath}}/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

// {{.BCName}}Service is a {{.SCName}} service.
type {{.BCName}}Service struct {
	v1.Unimplemented{{.BCName}}Server

	uc  *biz.{{.BCName}}Usecase
	log *log.Helper
}

// New{{.BCName}}Service new a {{.SCName}} service.
func New{{.BCName}}Service(uc *biz.{{.BCName}}Usecase, logger log.Logger) *{{.BCName}}Service {
	return &{{.BCName}}Service{uc: uc, log: log.NewHelper(logger)}
}

func (s *{{.BCName}}Service) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	s.log.WithContext(ctx).Infof("SayHello Received: %v", in.GetName())

	if in.GetName() == "error" {
		return nil, errors.New(fmt.Sprintf("user not found: %s", in.GetName()))
	}
	return &v1.HelloReply{Message: "Hello " + in.GetName()}, nil
}
`
)
