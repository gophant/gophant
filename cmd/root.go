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

var yes bool

func init() {
	// Global flags
	rootCmd.PersistentFlags().BoolVarP(&yes, "yes", "y", false, "Answer yes to prompts (non-interactive)")
	// Subcommands are added in separate files
}

// git push origin main && git tag v0.2.0 && git push origin v0.2.0 && go clean -modcache && go install github.com/gophant/gophant@latest
