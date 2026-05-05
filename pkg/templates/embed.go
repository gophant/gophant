package templates

import (
	"embed"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

//go:embed default/*
var defaultFS embed.FS

//go:embed mvc/**
var mvcFS embed.FS

//go:embed ddd/**
var dddFS embed.FS

// Embedded default TemplateSet (fallback)
var Embedded TemplateSet

func init() {
	if ts, err := loadTemplateSetFromFS(defaultFS, "default"); err == nil {
		Embedded = ts
	}
}

func loadTemplateSetFromFS(fsys fs.FS, base string) (TemplateSet, error) {
	read := func(name string) (string, error) {
		p := filepath.ToSlash(filepath.Join(base, name))
		b, err := fs.ReadFile(fsys, p)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}

	var ts TemplateSet
	var errAcc error
	if ts.GoMod, errAcc = read("go.mod.tmpl"); errAcc != nil {
		return TemplateSet{}, errAcc
	}
	if ts.Main, errAcc = read("main.go.tmpl"); errAcc != nil {
		return TemplateSet{}, errAcc
	}
	if ts.Env, errAcc = read(".env.example.tmpl"); errAcc != nil {
		return TemplateSet{}, errAcc
	}
	if ts.Readme, errAcc = read("README.md.tmpl"); errAcc != nil {
		return TemplateSet{}, errAcc
	}
	if ts.Config, errAcc = read("pkg_config_config.go.tmpl"); errAcc != nil {
		return TemplateSet{}, errAcc
	}
	if ts.Routes, errAcc = read("app_routes_routes.go.tmpl"); errAcc != nil {
		return TemplateSet{}, errAcc
	}
	if ts.Root, errAcc = read("cmd_root_root.go.tmpl"); errAcc != nil {
		return TemplateSet{}, errAcc
	}
	if ts.Serve, errAcc = read("cmd_serve_serve.go.tmpl"); errAcc != nil {
		return TemplateSet{}, errAcc
	}
	if ts.Gitignore, errAcc = read("gitignore.tmpl"); errAcc != nil {
		return TemplateSet{}, errAcc
	}
	return ts, nil
}

// LoadEmbeddedPath tries to load a template set from embedded mvc/ddd folders.
// Accepts paths like "mvc/react-app" or "./templates/mvc/react-app".
func LoadEmbeddedPath(path string) (TemplateSet, error) {
	p := strings.TrimPrefix(path, "./templates/")
	p = strings.TrimPrefix(p, "./")
	p = strings.TrimPrefix(p, "/")
	parts := strings.Split(p, string(filepath.Separator))
	if len(parts) < 2 {
		return TemplateSet{}, fmt.Errorf("invalid embedded template path: %s", path)
	}
	arch := parts[0]
	name := parts[1]
	switch arch {
	case "mvc":
		return loadTemplateSetFromFS(mvcFS, filepath.ToSlash(filepath.Join("mvc", name)))
	case "ddd":
		return loadTemplateSetFromFS(dddFS, filepath.ToSlash(filepath.Join("ddd", name)))
	case "default":
		return loadTemplateSetFromFS(defaultFS, filepath.ToSlash(filepath.Join("default", name)))
	default:
		return TemplateSet{}, fmt.Errorf("unknown architecture: %s", arch)
	}
}

// TemplateSet holds templates
type TemplateSet struct {
	GoMod     string
	Main      string
	Env       string
	Readme    string
	Config    string
	Routes    string
	Root      string
	Serve     string
	Gitignore string
}
