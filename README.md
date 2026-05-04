# Gophant — Go project scaffolder

Gophant is a small CLI that scaffolds a Go web project (Gin + GORM) with a convention-based structure so you can quickly start new apps.

## Install (recommended)

Install directly with Go (no manual build/move required):

  go install github.com/YOUR_GITHUB_USERNAME/gophant@latest

Notes:
- Ensure $(go env GOPATH)/bin or $GOBIN is in your PATH. Example: export PATH="$(go env GOPATH)/bin:$PATH"
- If installing from a tagged release use the tag: @v0.1.0

## Usage

Create a new project in the current working directory:

  gophant create myapp

Flags:
  -f, --force           Overwrite existing target directory if present
  -t, --templates PATH  Use custom templates from PATH (overrides embedded templates)

After creation:
  cd myapp
  go mod tidy
  cp .env.example .env
  go run cmd/main.go

## Development

Run tests and formatting locally:

  go test ./...
  go fmt ./...

Project structure and templates are embedded by default; use --templates to provide custom templates during development.

## CI & Releases

This repository uses GitHub Actions to run tests, vet, and formatting checks. For releases, use goreleaser to publish cross-compiled binaries and checksums.

## Contributing

Issues and PRs welcome. Keep changes small and add tests for generator behavior.

## License

MIT
