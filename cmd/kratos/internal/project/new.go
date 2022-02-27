package project

import (
	"context"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"github.com/go-kratos/kratos/cmd/kratos/v2/internal/base"
	tmpl "github.com/go-kratos/kratos/cmd/kratos/v2/internal/project/template"
)

// Project is a project template.
type Project struct {
	Name      string
	Path      string
	FullPath  string
	Namespace string
}

type Config struct {
	SCName  string
	BCName  string
	ModPath string
	Files   []file
}

type file struct {
	Path string
	Tmpl string
}

// New new a project from remote repo.
func (p *Project) New(ctx context.Context, dir string, layout string, branch string) error {
	to := path.Join(dir, p.Name)
	if _, err := os.Stat(to); !os.IsNotExist(err) {
		fmt.Printf("🚫 %s already exists\n", p.Name)
		override := false
		prompt := &survey.Confirm{
			Message: "📂 Do you want to override the folder ?",
			Help:    "Delete the existing folder and create the project.",
		}
		e := survey.AskOne(prompt, &override)
		if e != nil {
			return e
		}
		if !override {
			return err
		}
		os.RemoveAll(to)
	}
	fmt.Printf("🚀 Creating service %s, layout repo is %s, please wait a moment.\n\n", p.Name, layout)
	repo := base.NewRepo(layout, branch)
	if err := repo.CopyTo(ctx, to, p.Path, []string{".git", ".github"}); err != nil {
		return err
	}
	e := os.Rename(
		path.Join(to, "cmd", "server"),
		path.Join(to, "cmd", p.Namespace),
	)
	if e != nil {
		return e
	}

	e = p.deleteFiles()
	if e != nil {
		return e
	}

	e = p.templateReplace()
	if e != nil {
		return e
	}

	base.Tree(to, dir)

	fmt.Printf("\n🍺 Project creation succeeded %s\n", color.GreenString(p.Name))
	fmt.Print("💻 Use the following command to start the project 👇:\n\n")

	fmt.Println(color.WhiteString("$ cd %s", p.Name))
	fmt.Println(color.WhiteString("$ go generate ./..."))
	fmt.Println(color.WhiteString("$ go build -o ./bin/ ./... "))
	fmt.Println(color.WhiteString("$ ./bin/%s -conf ./configs\n", p.Name))
	fmt.Println("			🤝 Thanks for using Kratos")
	fmt.Println("	📚 Tutorial: https://go-kratos.dev/docs/getting-started/start")
	return nil
}

func (p *Project) templateReplace() error {
	config := Config{
		SCName: base.SmallCamel(p.Namespace),
		BCName: base.BigCamel(p.Namespace),
		Files: []file{
			{"/api/" + p.Namespace + "/" + p.Namespace + ".proto", tmpl.ApiProto},
			{"/cmd/" + p.Namespace + "/main.go", tmpl.Main},
			{"/cmd/" + p.Namespace + "/wire.go", tmpl.Wire},
			{"/cmd/" + p.Namespace + "/wire_gen.go", tmpl.WireGen},
			{"/internal/server/server.go", tmpl.Server},
			{"/internal/server/http.go/", tmpl.HTTPServer},
			{"/internal/server/grpc.go/", tmpl.GRPCServer},
			{"/internal/service/service.go", tmpl.Service},
			{"/internal/service/" + p.Namespace + ".go", tmpl.CustomService},
			{"/internal/biz/biz.go", tmpl.Biz},
			{"/internal/biz/" + p.Namespace + ".go", tmpl.CustomBiz},
			{"/internal/data/data.go", tmpl.Data},
			{"/internal/data/" + p.Namespace + ".go", tmpl.CustomData},
		},
	}
	mod, e := base.ModulePath(path.Join(p.FullPath, "go.mod"))
	if e != nil {
		return e
	}
	config.ModPath = mod

	for _, file := range config.Files {
		f := filepath.Join(p.FullPath, file.Path)
		dir := filepath.Dir(f)

		if _, e := os.Stat(dir); os.IsNotExist(e) {
			if e := os.MkdirAll(dir, 0755); e != nil {
				return e
			}
		}
		if e := p.write(f, file.Tmpl, config); e != nil {
			return e
		}
	}
	return nil
}

func (p *Project) deleteFiles() error {
	files := []string{
		filepath.Join(p.FullPath, "api/helloworld"),
		filepath.Join(p.FullPath, "internal/biz/greeter.go"),
		filepath.Join(p.FullPath, "internal/data/greeter.go"),
		filepath.Join(p.FullPath, "internal/service/greeter.go"),
	}
	for _, file := range files {
		if e := os.RemoveAll(file); e != nil {
			return e
		}
	}
	return nil
}

func (p *Project) write(file, tmpl string, config Config) error {
	fn := template.FuncMap{
		"title": strings.Title,
	}
	f, e := os.Create(file)
	if e != nil {
		return e
	}
	defer f.Close()

	t, e := template.New("f").Funcs(fn).Parse(tmpl)
	if e != nil {
		return e
	}
	return t.Execute(f, config)
}
