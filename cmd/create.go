package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/gophant/gophant/pkg/generator"
	"github.com/spf13/cobra"
)

var force bool

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [template] <app_name>",
	Short: "Create a new Gophant project from a template",
	Long: `Create a new Gin + GORM project with a convention-based structure.

Examples:
  gophant create myapp                  # use default embedded template
  gophant create react-app myapp        # use 'react-app' template from ./templates/react-app
  gophant create myapp --templates ./custom_templates

After creation:
  cd myapp
  go mod tidy
  go run cmd/main.go`,
	Args: cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		var templateName string
		var appName string
		if len(args) == 1 {
			appName = args[0]
		} else {
			templateName = args[0]
			appName = args[1]
		}

		// Determine templates dir to use
		templatesDirToUse := templatesDir
		if templatesDirToUse == "" && templateName != "" {
			// Look for ./templates/<templateName>
			templatesDirToUse = filepath.Join("./templates", templateName)
		}

		// Create the project
		fmt.Printf("🐘🐹 Creating Gophant project: %s\n", appName)
		if err := generator.CreateProject(appName, force, templatesDirToUse, yes); err != nil {
			return fmt.Errorf("failed to create project: %w", err)
		}

		// Success message
		fmt.Printf("\n✅ Gophant project created successfully!\n\n")
		fmt.Println("Next steps:")
		fmt.Printf("  cd %s\n", appName)
		fmt.Println("  go mod tidy")
		fmt.Println("  go run cmd/main.go")
		fmt.Println("\nThen visit: http://localhost:8080")

		return nil
	},
}

func init() {
	createCmd.Flags().BoolVarP(&force, "force", "f", false, "Overwrite existing directory if present")
	rootCmd.AddCommand(createCmd)
}
