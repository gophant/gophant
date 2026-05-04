package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gophant",
	Short: "Scaffolder for web developers learning Go",
	Long: `Gophant is a CLI tool that generates production-ready Gin + GORM
projects with sensible, convention-based defaults.

Perfect for developers transitioning to Go.`,
	Version: "0.1.0",
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Global flags would go here
	// For now, we'll add subcommands in separate files
}
