package cmd

import (
	"fmt"

	"github.com/gophant/gophant/pkg/generator"
	"github.com/spf13/cobra"
)

var force bool

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [app_name]",
	Short: "Create a new Gophant project",
	Long: `Create a new Gin + GORM project with a convention-based structure.

Example:
  gophant create myapp
  cd myapp
  go mod tidy
  go run cmd/main.go`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		appName := args[0]

		// Create the project
		fmt.Printf("🐘🐹 Creating Gophant project: %s\n", appName)
		if err := generator.CreateProject(appName, force, templatesDir, yes); err != nil {
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
