package generator

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/gophant/gophant/pkg/templates"
)

// ProjectData holds information about the project
type ProjectData struct {
	AppName string
	Year    string
}

// CreateProject creates a new Gophant project. If force is true, existing target will be removed.
// templatesDir overrides embedded templates when provided (path to templates directory).
func CreateProject(appName string, force bool, templatesDir string, assumeYes bool) error {
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
		filepath.Join(tempProjectPath, "app", "models"),
		filepath.Join(tempProjectPath, "app", "handlers"),
		filepath.Join(tempProjectPath, "app", "routes"),
		filepath.Join(tempProjectPath, "app", "middleware"),
		filepath.Join(tempProjectPath, "app", "database"),
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

	// Step 2: Load templates and create project files inside temp project
	projectData := ProjectData{
		AppName: appName,
		Year:    "2026",
	}

	ts, err := templates.Load(templatesDir)
	if err != nil {
		return fmt.Errorf("failed to load templates: %w", err)
	}

	files := map[string]string{
		filepath.Join(tempProjectPath, "go.mod"):                     ts.GoMod,
		filepath.Join(tempProjectPath, "main.go"):                    ts.Main,
		filepath.Join(tempProjectPath, ".env.example"):               ts.Env,
		filepath.Join(tempProjectPath, "README.md"):                  ts.Readme,
		filepath.Join(tempProjectPath, "pkg", "config", "config.go"): ts.Config,
		filepath.Join(tempProjectPath, "app", "routes", "routes.go"): ts.Routes,
		filepath.Join(tempProjectPath, "cmd", "root.go"):             ts.Root,
		filepath.Join(tempProjectPath, "cmd", "serve.go"):            ts.Serve,
		filepath.Join(tempProjectPath, ".gitignore"):                 ts.Gitignore,
	}

	fmt.Println("📝 Creating files...")
	for filePath, templateStr := range files {
		if err := createFileFromTemplate(filePath, templateStr, projectData); err != nil {
			return fmt.Errorf("failed to create file %s: %w", filePath, err)
		}
	}

	// Step 3: Move temp project to target location atomically
	if _, err := os.Stat(appName); err == nil {
		// target exists
		if force {
			// remove existing when force is true
			if err := os.RemoveAll(appName); err != nil {
				return fmt.Errorf("failed to remove existing directory %s: %w", appName, err)
			}
		} else {
			// If user supplied assumeYes (-y), proceed without prompt
			if assumeYes {
				if err := os.RemoveAll(appName); err != nil {
					return fmt.Errorf("failed to remove existing directory %s: %w", appName, err)
				}
			} else {
				// interactive confirmation when possible
				if fi, _ := os.Stdin.Stat(); (fi.Mode() & os.ModeCharDevice) != 0 {
					// prompt
					reader := bufio.NewReader(os.Stdin)
					fmt.Printf("directory '%s' already exists. Overwrite? (y/N): ", appName)
					ans, _ := reader.ReadString('\n')
					ans = strings.TrimSpace(ans)
					if strings.EqualFold(ans, "y") || strings.EqualFold(ans, "yes") {
						if err := os.RemoveAll(appName); err != nil {
							return fmt.Errorf("failed to remove existing directory %s: %w", appName, err)
						}
					} else {
						return fmt.Errorf("aborted: directory '%s' exists", appName)
					}
				} else {
					// non-interactive: require --force
					return fmt.Errorf("directory '%s' already exists (use --force to overwrite)", appName)
				}
			}
		}
	}

	if err := moveDir(tempProjectPath, appName); err != nil {
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

func moveDir(src, dst string) error {
	// Try fast move first
	err := os.Rename(src, dst)
	if err == nil {
		return nil
	}

	// Fallback: copy + delete (works across disks)
	if err := copyDir(src, dst); err != nil {
		return err
	}

	return os.RemoveAll(src)
}

func copyDir(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		targetPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			return os.MkdirAll(targetPath, info.Mode())
		}

		return copyFile(path, targetPath, info.Mode())
	})
}

func copyFile(src, dst string, perm os.FileMode) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err = io.Copy(out, in); err != nil {
		return err
	}

	return os.Chmod(dst, perm)
}
