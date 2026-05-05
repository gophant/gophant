Gophant — Go project scaffolder

Gophant scaffolds Go web projects (Gin + GORM conventions) with a consistent, convention-based layout so teams start from the same structure and configuration.

This repository contains the CLI generator (gophant) and several built-in templates/skeletons:
- pkg/templates/default — the built-in default templated files used when no override is provided
- pkg/templates/mvc — an MVC starter skeleton
- pkg/templates/ddd — a DDD starter skeleton

Goals
- Fast project bootstrap with consistent layout and sensible defaults
- Support embedded built-in skeletons and local/custom templates
- Simple, scriptable CLI flags for architecture selection
- Template files are text/template-based so values are substituted at generation time

Install

  go install github.com/gophant/gophant@latest

Make sure your Go bin is in PATH:

  export PATH="$(go env GOPATH)/bin:$PATH"

Quick usage

  # Create a project using the embedded default template
  gophant create myapp

  # Create using embedded MVC skeleton
  gophant create myapp -a mvc

  # Create using embedded DDD skeleton
  gophant create myapp -a ddd

  # Create using embedded MVC skeleton
  gophant create myapp -a mvc

  # Create using embedded DDD skeleton
  gophant create myapp -a ddd

  # List available embedded templates
  gophant templates

CLI flags (summary)
- -a, --arch, --architecture string
  - Select architecture group. Valid values: `mvc`, `ddd`, `default` (optional).
  - If omitted, the embedded `default` template is used.
- --template string
  - Variant name under the chosen architecture. When provided, gophant first looks for a local skeleton at `./templates/<arch>/<variant>`; if found it is used. (Note: do not use slash-style template names as positional args.)
- -t, --templates string
  - Use a custom templates directory (local path). When set, it overrides embedded templates.
- -f, --force
  - Overwrite an existing target directory.
- templates (command)
  - `gophant templates` prints embedded architectures and variants and lists any local `./templates` dirs.

Behavior & selection rules
1. If `--templates` (local path) is given and it exists as a directory, that skeleton is used.
2. Otherwise, if `-a/--arch` is provided:
   - If `--template` is provided, gophant attempts to use `./templates/<arch>/<template>` (local) first. If that local path does not exist, the generator attempts to use the embedded architecture skeleton (for `arch`).
   - If `--template` is not provided, gophant attempts to use embedded `arch` skeleton.
3. If no arch and no templates flag, embedded `default` is used.
4. The CLI rejects slash-style positional template arguments (e.g. `gophant create mvc/react-app myapp`). Use flags instead.

Template files and conventions
- Template files use Go text/template. Common conventions:
  - go.mod.tmpl — module, should use `{{.AppName}}` for module path
  - *.go.tmpl — Go source templates (will be formatted after generation)
  - .env.example, README.md, migrations, resources — copy as-is or as templates
- Template data available:
  - {{.AppName}} — application name (module path)
  - {{.Year}} — year (generator populates a value)
- Local template layout: a full skeleton directory mirrors the final project layout (e.g., put main.go.tmpl at root, cmd/, pkg/, migrations/, etc.). Any file with `.tmpl` will be processed as a text/template (and Go files will be gofmt-ed).

Commands & examples (detailed)
- Create default app:
  gophant create todoapp
  cd todoapp
  cp .env.example .env
  go mod tidy
  go run main.go

- Create with arch:
  gophant create blog -a mvc

- Create with arch + local variant:
  # local folder ./templates/mvc/react-app exists
  gophant create blog -a mvc --template react-app

- Use custom templates directory:
  gophant create svc --templates /home/me/my-templates

Developer notes
- Embedded templates are located in pkg/templates and are embedded at build time.
- The loader prefers local templates (./templates or --templates) over embedded skeletons.
- To add a new built-in template, add files under pkg/templates/<arch> (use .tmpl suffix for templated files). Ensure required top-level template files (go.mod.tmpl, main.go.tmpl, pkg_config_config.go.tmpl, app_routes_routes.go.tmpl, cmd_root_root.go.tmpl, cmd_serve_serve.go.tmpl, gitignore.tmpl) are present if you want the embedded arch to be directly selectable.
- Tests:
  - Run tests: `go test ./...` or per-package `go test ./pkg/templates`
  - There are unit tests verifying template enumeration and loader behavior.

CI & release suggestions
- Add a GitHub Actions workflow to run:
  - go fmt ./...
  - go test ./...
  - go build ./...
- Use goreleaser for binary releases.

Notes about safety and UX
- The CLI prints what it’s using when creating (architecture and template messages).
- The generator tries a best-effort `git init` in the created project but it’s non-fatal if git isn't available.

Contributing
- Open issues for feature requests.
- Submit small, focused PRs and include tests for behavior changes.
- Keep templates minimal and document placeholders in README-TEMPLATE.md inside template skeletons.

License
- MIT
