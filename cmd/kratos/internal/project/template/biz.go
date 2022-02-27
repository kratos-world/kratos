package template

var (
	Biz = `package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(New{{.BCName}}Usecase)`
	
	CustomBiz = `package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type {{.BCName}} struct {
	Hello string
}

type {{.BCName}}Repo interface {
	Create{{.BCName}}(context.Context, *{{.BCName}}) error
	Update{{.BCName}}(context.Context, *{{.BCName}}) error
}

type {{.BCName}}Usecase struct {
	repo {{.BCName}}Repo
	log  *log.Helper
}

func New{{.BCName}}Usecase(repo {{.BCName}}Repo, logger log.Logger) *{{.BCName}}Usecase {
	return &{{.BCName}}Usecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *{{.BCName}}Usecase) Create(ctx context.Context, g *{{.BCName}}) error {
	return uc.repo.Create{{.BCName}}(ctx, g)
}

func (uc *{{.BCName}}Usecase) Update(ctx context.Context, g *{{.BCName}}) error {
	return uc.repo.Update{{.BCName}}(ctx, g)
}
`
)
