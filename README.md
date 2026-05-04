# Gophant — Go project scaffolder

Gophant is a small CLI that scaffolds a Go web project (Gin + GORM) with a convention-based structure so you can quickly start new apps.

## Install (local)

- Build and move to a directory in PATH:
  go build -o gophant .
  sudo mv gophant /usr/local/bin

- Or install directly (recommended for releases):
  go install github.com/gophant/cli@latest

## Usage

Create a new project in the current working directory:

  gophant create myapp

Flags:
  -f, --force   Overwrite existing target directory if present

After creation:
  cd myapp
  go mod tidy
  cp .env.example .env
  go run cmd/main.go

## Development

Run tests and formatting locally:

  go test ./...
  go fmt ./...

Project structure and templates are embedded so the generator is easy to maintain.

## CI & Releases

This repository uses GitHub Actions to run tests, vet, and formatting checks. Releases should publish binaries (use goreleaser for cross-compilation and artifacts).

## Contributing

Issues and PRs welcome. Keep changes small and add tests for generator behavior.

## License

MIT

