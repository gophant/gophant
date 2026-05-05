Gophant — Go project scaffolder

Gophant scaffolds Go web projects (Gin + GORM conventions) with a consistent, convention-based layout so teams start from the same structure and configuration.

This repository contains the CLI generator (gophant) and several built-in templates/skeletons:
- pkg/templates/default — the built-in default templated files used when no override is provided
- pkg/templates/mvc — an MVC starter skeleton
- pkg/templates/ddd — a DDD starter skeleton

Goals
- Fast project bootstrap with consistent layout and sensible defaults
- Keep the CLI minimal and easy to use — choose architecture with a single flag
- Templates are text/template-based so values are substituted at generation time

Install

  go install github.com/gophant/gophant@latest

Make sure your Go bin is in PATH:

  export PATH="$GOPATH/bin:$PATH"

Quick usage

  # Create a project using the embedded default template
  gophant create myapp

  # Create using embedded MVC skeleton
  gophant create myapp -a mvc

  # Create using embedded DDD skeleton
  gophant create myapp -a ddd

  # List available embedded templates (see below for lists)
  # Note: the interactive templates command was removed to keep the CLI minimal — details are in this README

Template file structures (what's embedded now)

pkg/templates/default (primary files — top-level templates)
- internal_routes_routes.go.tmpl
- main.go.tmpl
- README.md.tmpl
- gitignore.tmpl
- .env.example.tmpl
- cmd_root_root.go.tmpl
- go.mod.tmpl
- cmd_serve_serve.go.tmpl
- pkg_config_config.go.tmpl
- app_routes_routes.go.tmpl

pkg/templates/mvc (example important files up to two levels)
- config/.gitkeep
- config/placeholder.txt
- .gophant.yml
- README.md
- resources/.gitkeep
- resources/placeholder.txt
- tests/.gitkeep
- tests/placeholder.txt
- internal/.gitkeep
- internal/placeholder.txt
- README-TEMPLATE.md
- storage/.gitkeep
- storage/placeholder.txt
- pkg/.gitkeep
- pkg/placeholder.txt
- .gophant/.gitkeep
- .gophant/placeholder.txt
- go.mod.tmpl
- routes/routes.go.tmpl
- .env.example
- cmd/.gitkeep
- cmd/placeholder.txt
- database/.gitkeep
- database/placeholder.txt
- app/.gitkeep
- app/placeholder.txt

pkg/templates/ddd (example important files up to two levels)
- application/.gitkeep
- application/placeholder.txt
- .gophant.yml
- README.md
- resources/.gitkeep
- resources/placeholder.txt
- infrastructure/.gitkeep
- infrastructure/placeholder.txt
- domain/.gitkeep
- domain/placeholder.txt
- tests/.gitkeep
- tests/placeholder.txt
- README-TEMPLATE.md
- storage/.gitkeep
- storage/placeholder.txt
- pkg/.gitkeep
- pkg/placeholder.txt
- go.mod.tmpl
- routes/routes.go.tmpl
- .env.example
- cmd/.gitkeep
- cmd/placeholder.txt
- database/.gitkeep
- database/placeholder.txt
- interfaces/.gitkeep
- interfaces/placeholder.txt

CLI flags (summary)
- -a, --arch, --architecture string
  - Select architecture group. Valid values: `mvc`, `ddd`, `default` (optional).
  - If omitted, the embedded `default` template is used.
- -f, --force
  - Overwrite an existing target directory.
- -y, --yes
  - Answer yes to prompts (non-interactive)

Behavior & selection rules
1. The CLI is intentionally minimal: it uses embedded templates only.
2. If `-a/--arch` is provided, gophant will use the embedded skeleton for that architecture.
3. If no arch provided, embedded `default` is used.
4. Local/custom templates support was removed to keep UX simple. Future versions may reintroduce it behind an opt-in flag.

Template files and conventions
- Template files use Go text/template. Common conventions:
  - go.mod.tmpl — module, should use `{{.AppName}}` for module path
  - *.go.tmpl — Go source templates (will be formatted after generation)
  - .env.example, README.md, migrations, resources — copy as-is or as templates
- Template data available:
  - {{.AppName}} — application name (module path)
  - {{.Year}} — year (generator populates a value)

Commands & examples (detailed)
- Create default app:
  gophant create todoapp
  cd todoapp
  cp .env.example .env
  go mod tidy
  go run main.go

- Create with arch:
  gophant create blog -a mvc

Developer notes
- Embedded templates are located in pkg/templates and are embedded at build time.
- To add a new built-in template, add files under pkg/templates/<arch> (use .tmpl suffix for templated files). Ensure required top-level template files are present if you want the embedded arch to be directly selectable.
- Tests:
  - Run tests: `go test ./...` or per-package `go test ./pkg/templates`

CI & release suggestions
- Add a GitHub Actions workflow to run:
  - go fmt ./...
  - go test ./...
  - go build ./...
- Use goreleaser for binary releases.

Contributing
- Open issues for feature requests.
- Submit small, focused PRs and include tests for behavior changes.

License
- MIT
