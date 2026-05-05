package templates

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
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
// Accepts paths like "mvc", "mvc/react-app" or "./templates/mvc".
func LoadEmbeddedPath(path string) (TemplateSet, error) {
	p := strings.TrimPrefix(path, "./templates/")
	p = strings.TrimPrefix(p, "./")
	p = strings.TrimPrefix(p, "/")
	// normalize to forward slashes
	p = filepath.ToSlash(p)
	parts := strings.Split(p, "/")
	if len(parts) == 0 || parts[0] == "" {
		return TemplateSet{}, fmt.Errorf("invalid embedded template path: %s", path)
	}
	arch := parts[0]
	rest := ""
	if len(parts) > 1 {
		rest = strings.Join(parts[1:], "/")
	}
	switch arch {
	case "mvc":
		if rest == "" {
			return loadTemplateSetFromFS(mvcFS, "mvc")
		}
		return loadTemplateSetFromFS(mvcFS, filepath.ToSlash(filepath.Join("mvc", rest)))
	case "ddd":
		if rest == "" {
			return loadTemplateSetFromFS(dddFS, "ddd")
		}
		return loadTemplateSetFromFS(dddFS, filepath.ToSlash(filepath.Join("ddd", rest)))
	case "default":
		if rest == "" {
			return loadTemplateSetFromFS(defaultFS, "default")
		}
		return loadTemplateSetFromFS(defaultFS, filepath.ToSlash(filepath.Join("default", rest)))
	default:
		return TemplateSet{}, fmt.Errorf("unknown architecture: %s", arch)
	}
}

// CopyEmbeddedSkeleton copies files from an embedded skeleton (e.g., "mvc" or "ddd") to dest.
// It returns an error if the embedded path does not exist or is empty.
func CopyEmbeddedSkeleton(path string, dest string) error {
	p := strings.TrimPrefix(path, "./templates/")
	p = strings.TrimPrefix(p, "./")
	p = strings.TrimPrefix(p, "/")
	p = filepath.ToSlash(p)
	parts := strings.Split(p, "/")
	if len(parts) == 0 || parts[0] == "" {
		return fmt.Errorf("invalid embedded template path: %s", path)
	}
	arch := parts[0]
	rest := ""
	if len(parts) > 1 {
		rest = strings.Join(parts[1:], "/")
	}

	var fsys fs.FS
	var base string
	switch arch {
	case "mvc":
		fsys = mvcFS
		base = "mvc"
	case "ddd":
		fsys = dddFS
		base = "ddd"
	case "default":
		fsys = defaultFS
		base = "default"
	default:
		return fmt.Errorf("unknown architecture: %s", arch)
	}

	if rest != "" {
		base = filepath.ToSlash(filepath.Join(base, rest))
	}

	// Walk embedded FS and write files into dest preserving relative paths
	return fs.WalkDir(fsys, base, func(pth string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(base, pth)
		if err != nil {
			return err
		}
		target := filepath.Join(dest, rel)
		if d.IsDir() {
			return os.MkdirAll(target, 0o755)
		}
		// Read file from embedded FS
		b, err := fs.ReadFile(fsys, pth)
		if err != nil {
			return err
		}
		// Ensure parent dir
		if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
			return err
		}
		// Write file
		if err := os.WriteFile(target, b, 0o644); err != nil {
			return err
		}
		return nil
	})
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
