Gophant — Go project scaffolder

Gophant is a simple CLI that scaffolds Go web projects (Gin + GORM) with a clear, convention-based layout.
It provides embedded templates for a "default" project and starter skeletons for "mvc" and "ddd" architectures.

Quick install

  go install github.com/gophant/gophant@latest

Make sure your Go bin is in PATH:

  export PATH="$(go env GOPATH)/bin:$PATH"

Usage

  gophant create [template] <app_name>

Examples

  # Create a new project using embedded default template
  gophant create myapp

  # Create using embedded mvc skeleton (top-level name)
  gophant create mvc myapp

  # Create using embedded mvc skeleton
  gophant create mvc myapp

  # Create using embedded ddd skeleton
  gophant create ddd myapp

  # Use a custom templates directory on disk (overrides embedded)
  gophant create myapp --templates ./my-templates

  # Force overwrite if target exists
  gophant create myapp --force

Post-create quick start

  cd myapp
  cp .env.example .env
  go mod tidy
  go run main.go

How templates work

- Built-in templates are embedded under pkg/templates:
  - pkg/templates/default  : the default templated files used when no override is provided
  - pkg/templates/mvc      : mvc skeleton files (ready as a starting point)
  - pkg/templates/ddd      : ddd skeleton files (starter structure)

- Template file naming conventions
  Template files use Go's text/template syntax and the following file name conventions inside a template skeleton:
    - go.mod.tmpl           (module file; uses {{.AppName}} placeholder)
    - *.go.tmpl             (Go source files treated as templates)
    - .env.example, README.md, migrations, resources, etc.

- Template data available to templates
  - {{.AppName}}  — name of the generated app (used to set module path in go.mod)
  - {{.Year}}     — year value provided by the generator

Using custom templates

1. Create a directory that mirrors the layout you want the final project to have.
2. Use .tmpl suffix for files that require placeholder substitution (e.g., main.go.tmpl, go.mod.tmpl).
3. Run:
     gophant create myapp --templates ./path/to/your/templates

Advanced: embedded template selection

- If the CLI argument for template matches an embedded path like "mvc" or "mvc/react-app",
  Gophant will prefer the embedded skeleton when no local templates directory is provided.
- You can also place a templates/ directory in your current working directory and the CLI will pick it up automatically.

Project layout produced

A generated project will look like this (example):

  myapp/
  ├── cmd/
  ├── app/ or internal/
  ├── pkg/
  ├── migrations/
  ├── .env.example
  ├── go.mod
  └── main.go

Validation & release notes

- The repository builds successfully (go build ./...).
- Templates are embedded and validated for compilation-time embedding (placeholder files added where necessary).

Contributing

- Open an issue to discuss changes or submit a PR.
- Prefer small, focused changes and include tests for behavior changes.

CI & publishing

- Add a GitHub Actions workflow to run `go test ./...` and `gofmt` on each push.
- Use goreleaser for publishing binaries when creating a release tag.

License

MIT
