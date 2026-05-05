package templates

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// Load loads templates from overrideDir if provided or from common locations, otherwise returns embedded templates.
// If overrideDir points to a non-existent local path but matches an embedded template path
// like "mvc/react-app" or "./templates/mvc/react-app", the embedded template will be used.
func Load(overrideDir string) (TemplateSet, error) {
	// If explicit override provided
	if overrideDir != "" {
		// If it's an actual directory on disk, prefer it
		if ts, err := loadFromDir(overrideDir); err == nil {
			return ts, nil
		}

		// Try loading as embedded skeleton (supports paths like "mvc/react-app" or "./templates/mvc/react-app")
		if ts, err := LoadEmbeddedPath(overrideDir); err == nil {
			return ts, nil
		} else {
			return TemplateSet{}, fmt.Errorf("failed to load templates from %s: %w", overrideDir, err)
		}
	}

	// Check current working directory ./templates
	if ts, err := loadFromDir("./templates"); err == nil {
		return ts, nil
	}

	// Check $HOME/.gophant/templates
	home := os.Getenv("HOME")
	if home != "" {
		path := filepath.Join(home, ".gophant", "templates")
		if ts, err := loadFromDir(path); err == nil {
			return ts, nil
		}
	}

	// Fallback to embedded default
	return Embedded, nil
}

func loadFromDir(dir string) (TemplateSet, error) {
	info, err := os.Stat(dir)
	if err != nil || !info.IsDir() {
		return TemplateSet{}, fs.ErrNotExist
	}

	read := func(name string) (string, error) {
		b, err := os.ReadFile(filepath.Join(dir, name))
		if err != nil {
			return "", err
		}
		return string(b), nil
	}

	var ts TemplateSet
	var errAcc error
	if ts.GoMod, err = read("go.mod.tmpl"); err != nil {
		errAcc = err
	}
	if ts.Main, err = read("main.go.tmpl"); err != nil {
		errAcc = err
	}
	if ts.Env, err = read(".env.example.tmpl"); err != nil {
		errAcc = err
	}
	if ts.Readme, err = read("README.md.tmpl"); err != nil {
		errAcc = err
	}
	if ts.Config, err = read("pkg_config_config.go.tmpl"); err != nil {
		errAcc = err
	}
	if ts.Routes, err = read("app_routes_routes.go.tmpl"); err != nil {
		errAcc = err
	}
	if ts.Root, err = read("cmd_root_root.go.tmpl"); err != nil {
		errAcc = err
	}
	if ts.Serve, err = read("cmd_serve_serve.go.tmpl"); err != nil {
		errAcc = err
	}
	if ts.Gitignore, err = read("gitignore.tmpl"); err != nil {
		errAcc = err
	}

	if errAcc != nil {
		return TemplateSet{}, errAcc
	}
	return ts, nil
}
