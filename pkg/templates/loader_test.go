package templates

import (
	"os"
	"path/filepath"
	"testing"
)

func writeFile(t *testing.T, dir, name, content string) {
	t.Helper()
	p := filepath.Join(dir, name)
	if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
		t.Fatalf("failed to mkdir for %s: %v", p, err)
	}
	if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
		t.Fatalf("failed to write %s: %v", p, err)
	}
}

func TestLoadFromDir(t *testing.T) {
	d := t.TempDir()
	// create required template files
	writeFile(t, d, "go.mod.tmpl", "module {{.AppName}}")
	writeFile(t, d, "main.go.tmpl", "package main")
	writeFile(t, d, ".env.example.tmpl", "ENV=VALUE")
	writeFile(t, d, "README.md.tmpl", "readme")
	writeFile(t, d, "pkg_config_config.go.tmpl", "package config")
	writeFile(t, d, "app_routes_routes.go.tmpl", "package routes")
	writeFile(t, d, "cmd_root_root.go.tmpl", "package cmd")
	writeFile(t, d, "cmd_serve_serve.go.tmpl", "package cmd")
	writeFile(t, d, "gitignore.tmpl", "node_modules/")

	ts, err := Load(d)
	if err != nil {
		t.Fatalf("Load returned error: %v", err)
	}

	if ts.GoMod != "module {{.AppName}}" {
		t.Fatalf("unexpected GoMod: %q", ts.GoMod)
	}
	if ts.Main != "package main" {
		t.Fatalf("unexpected Main: %q", ts.Main)
	}
	if ts.Env != "ENV=VALUE" {
		t.Fatalf("unexpected Env: %q", ts.Env)
	}
}

func TestLoadEmbeddedFallback(t *testing.T) {
	// Request a non-existent local dir but a known embedded arch identifier
	ts, err := Load("default")
	if err != nil {
		t.Fatalf("expected embedded default to load, got error: %v", err)
	}
	// At least one primary template should be present
	if ts.GoMod == "" && ts.Main == "" {
		t.Fatalf("expected embedded template to contain GoMod or Main, got empty")
	}
}
