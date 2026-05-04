package templates

import (
	"io/fs"
	"os"
	"path/filepath"
)

// Load loads templates from overrideDir if provided or from common locations, otherwise returns embedded templates.
func Load(overrideDir string) (TemplateSet, error) {
	// If explicit override provided
	if overrideDir != "" {
		if ts, err := loadFromDir(overrideDir); err == nil {
			return ts, nil
		} else {
			return TemplateSet{}, err
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

	// Fallback to embedded
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
	if ts.Routes, err = read("internal_routes_routes.go.tmpl"); err != nil {
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
