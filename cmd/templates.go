package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/gophant/gophant/pkg/templates"
	"github.com/spf13/cobra"
)

var templatesCmd = &cobra.Command{
	Use:   "templates",
	Short: "List available templates (embedded and local)",
	RunE: func(cmd *cobra.Command, args []string) error {
		m, err := templates.ListEmbeddedTemplates()
		if err != nil {
			return err
		}

		fmt.Println("Embedded templates:")
		for arch, variants := range m {
			fmt.Printf("- %s: %s\n", arch, strings.Join(variants, ", "))
		}

		// Show local templates if present
		if info, err := os.Stat("./templates"); err == nil && info.IsDir() {
			fmt.Println("\nLocal templates (./templates):")
			entries, err := os.ReadDir("./templates")
			if err == nil {
				for _, e := range entries {
					if e.IsDir() {
						fmt.Printf("- %s\n", e.Name())
					}
				}
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(templatesCmd)
}
