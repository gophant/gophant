package cmd

import (
	"fmt"
	"github.com/gophant/gophant/pkg/generator"
	"github.com/spf13/cobra"
)

var force bool
var arch string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create <app_name>",
	Short: "Create a new Gophant project from a template",
	Long: `Create a new Gin + GORM project with a convention-based structure.

Examples:
  gophant create myapp                  # use default embedded template
  gophant create myapp -a mvc           # use embedded mvc skeleton
  gophant create myapp -a mvc --template react-app  # use named variant under mvc (if available on disk)
  gophant create myapp --templates ./custom_templates

After creation:
  cd myapp
  go mod tidy
  go run main.go`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		appName := args[0]

		// Validate arch
		if arch != "" && arch != "mvc" && arch != "ddd" && arch != "default" {
			return fmt.Errorf("invalid arch: %s (must be 'mvc', 'ddd' or 'default')", arch)
		}

		// Determine templates dir to use: only embedded arch or default (no local templates support)
		templatesDirToUse := ""
		if arch != "" {
			// pass arch to loader so it may load embedded skeleton for the architecture
			templatesDirToUse = arch
		}

		// Create the project
		fmt.Printf("🐘🐹 Creating Gophant project: %s\n", appName)
		if arch != "" {
			fmt.Printf("Using architecture: %s\n", arch)
		}

		if err := generator.CreateProject(appName, force, templatesDirToUse, yes); err != nil {
			return fmt.Errorf("failed to create project: %w", err)
		}

		// Success message
		fmt.Printf("\n✅ Gophant project created successfully!\n\n")
		fmt.Println("Next steps:")
		fmt.Printf("  cd %s\n", appName)
		fmt.Println("  go mod tidy")
		fmt.Println("  go run main.go")
		fmt.Println("\nThen visit: http://localhost:8080")

		return nil
	},
}

func init() {
	createCmd.Flags().BoolVarP(&force, "force", "f", false, "Overwrite existing directory if present")
	createCmd.Flags().StringVarP(&arch, "arch", "a", "", "Architecture template group to use (mvc|ddd)")
	createCmd.Flags().StringVar(&arch, "architecture", "", "Architecture template group to use (mvc|ddd)")
	rootCmd.AddCommand(createCmd)
}
