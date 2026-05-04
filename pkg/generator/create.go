package generator

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"text/template"
)

// ProjectData holds information about the project
type ProjectData struct {
	AppName string
	Year    string
}

// CreateProject creates a new Gophant project. If force is true, existing target will be removed.
func CreateProject(appName string, force bool) error {
	// Validate app name
	if !isValidName(appName) {
		return fmt.Errorf("invalid app name: %s. Use letters, numbers, underscore or hyphen, and start with a letter", appName)
	}

	// Create temporary working directory
	tempDir, err := os.MkdirTemp("", "gophant-*")
	if err != nil {
		return fmt.Errorf("failed to create temp dir: %w", err)
	}
	defer os.RemoveAll(tempDir)

	tempProjectPath := filepath.Join(tempDir, appName)

	// Step 1: Create all required directories inside temp project
	dirs := []string{
		filepath.Join(tempProjectPath, "cmd"),
		filepath.Join(tempProjectPath, "internal", "models"),
		filepath.Join(tempProjectPath, "internal", "handlers"),
		filepath.Join(tempProjectPath, "internal", "routes"),
		filepath.Join(tempProjectPath, "internal", "middleware"),
		filepath.Join(tempProjectPath, "internal", "database"),
		filepath.Join(tempProjectPath, "pkg", "config"),
		filepath.Join(tempProjectPath, "pkg", "validators"),
		filepath.Join(tempProjectPath, "migrations"),
	}

	fmt.Println("📁 Creating directories...")
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	// Step 2: Create project files inside temp project
	projectData := ProjectData{
		AppName: appName,
		Year:    "2026",
	}

	files := map[string]string{
		filepath.Join(tempProjectPath, "go.mod"):                          GoModTemplate,
		filepath.Join(tempProjectPath, "cmd", "main.go"):                  MainTemplate,
		filepath.Join(tempProjectPath, ".env.example"):                    EnvTemplate,
		filepath.Join(tempProjectPath, "README.md"):                       ReadmeTemplate,
		filepath.Join(tempProjectPath, "pkg", "config", "config.go"):      ConfigTemplate,
		filepath.Join(tempProjectPath, "internal", "routes", "routes.go"): RoutesTemplate,
		filepath.Join(tempProjectPath, ".gitignore"):                      GitignoreTemplate,
	}

	fmt.Println("📝 Creating files...")
	for filePath, templateStr := range files {
		if err := createFileFromTemplate(filePath, templateStr, projectData); err != nil {
			return fmt.Errorf("failed to create file %s: %w", filePath, err)
		}
	}

	// Step 3: Move temp project to target location atomically
	if _, err := os.Stat(appName); err == nil {
		if !force {
			return fmt.Errorf("directory '%s' already exists", appName)
		}
		// remove existing when force is true
		if err := os.RemoveAll(appName); err != nil {
			return fmt.Errorf("failed to remove existing directory %s: %w", appName, err)
		}
	}

	if err := os.Rename(tempProjectPath, appName); err != nil {
		return fmt.Errorf("failed to move project into place: %w", err)
	}

	// Try to initialize git (best-effort)
	if err := tryGitInit(appName); err != nil {
		// Non-fatal; print warning
		fmt.Printf("warning: git init failed: %v\n", err)
	}

	return nil
}

// createFileFromTemplate creates a file from a template string. Go files are gofmt-ed.
func createFileFromTemplate(filePath string, templateStr string, data interface{}) error {
	tmpl, err := template.New("file").Parse(templateStr)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	// If this is a .go file, format it
	out := buf.Bytes()
	if filepath.Ext(filePath) == ".go" {
		if formatted, err := format.Source(out); err == nil {
			out = formatted
		}
	}

	// Ensure parent directory exists (templates may point at nested paths)
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		return fmt.Errorf("failed to create parent dir for %s: %w", filePath, err)
	}

	if err := os.WriteFile(filePath, out, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}

	return nil
}

func isValidName(name string) bool {
	if name == "" {
		return false
	}
	// reject path separators
	if name == "." || name == ".." || regexp.MustCompile(`[\\/]+`).MatchString(name) {
		return false
	}
	// must start with a letter and contain letters, numbers, underscore or hyphen
	re := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_-]*$`)
	return re.MatchString(name)
}

// tryGitInit runs `git init` in the newly created project dir. Best-effort only.
func tryGitInit(dir string) error {
	cmd := exec.Command("git", "init")
	cmd.Dir = dir
	// Discard output
	cmd.Stdout = io.Discard
	cmd.Stderr = io.Discard
	return cmd.Run()
}
