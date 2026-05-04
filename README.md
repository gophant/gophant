# Gophant — Go project scaffolder

Gophant is a simple CLI that creates a Go web project (Gin + GORM) with a clear, convention-based layout.

## Install (recommended)

Install with Go (one command):

  go install github.com/YOUR_GITHUB_USERNAME/gophant@latest

Make sure your Go bin is in PATH:

  export PATH="$(go env GOPATH)/bin:$PATH"

(If using a released version, replace @latest with a tag like @v0.1.0.)

## Usage

Create a new project:

  gophant create myapp

Flags:
  -f, --force           Overwrite existing target directory if present
  -t, --templates PATH  Use custom templates (see examples below)

Quick start after creation:

  cd myapp
  cp .env.example .env
  go mod tidy
  go run cmd/main.go

## Templates — examples (easy)

Default (built-in templates):

  gophant create myapp

Use templates from current directory:

1. Create a folder named "templates" where you run the command.
2. Put these files inside the folder:
   - go.mod.tmpl
   - main.go.tmpl
   - .env.example.tmpl
   - README.md.tmpl
   - pkg_config_config.go.tmpl
   - internal_routes_routes.go.tmpl
   - gitignore.tmpl
3. Run:

  gophant create myapp

Use templates from a specific path:

  gophant create myapp -t /home/me/custom-templates

The CLI will use your files instead of the built-in templates.

## Run tests & format code (local)

- Run tests:

  go test ./...

- Format code (keeps style consistent):

  go fmt ./...

Tests check behavior; fmt fixes code formatting.

## CI & Releases (simple)

- Enable GitHub Actions to run tests and go fmt on each push (helps catch problems early).
- Use goreleaser to build and publish binaries for different OS/architectures when you create a release tag.

## Contributing (simple)

- Open an issue to discuss changes or submit a small PR.
- Prefer small, focused PRs and add tests for any code changes.

## License

MIT
