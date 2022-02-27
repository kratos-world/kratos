package template

var (
	Data = `package data

import (
	"{{.ModPath}}/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, New{{.BCName}}Repo)

// Data .
type Data struct {
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}
`
	
	CustomData = `package data

import (
	"context"
	"{{.ModPath}}/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type {{.SCName}}Repo struct {
	data *Data
	log  *log.Helper
}

// New{{.BCName}}Repo .
func New{{.BCName}}Repo(data *Data, logger log.Logger) biz.{{.BCName}}Repo {
	return &{{.SCName}}Repo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *{{.SCName}}Repo) Create{{.BCName}}(ctx context.Context, g *biz.{{.BCName}}) error {
	return nil
}

func (r *{{.SCName}}Repo) Update{{.BCName}}(ctx context.Context, g *biz.{{.BCName}}) error {
	return nil
}
`
)
