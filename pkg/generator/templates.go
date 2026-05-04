package generator

// GoModTemplate is the go.mod template
var GoModTemplate = `module {{.AppName}}

go 1.21

require (
	github.com/gin-gonic/gin v1.9.1
	gorm.io/gorm v1.25.5
	gorm.io/driver/mysql v1.5.2
	github.com/joho/godotenv v1.5.1
)
`

// MainTemplate is the cmd/main.go template
var MainTemplate = `package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"{{.AppName}}/internal/routes"
	"{{.AppName}}/pkg/config"
)

func main() {
	// Load environment variables
	godotenv.Load()

	// Initialize Gin router
	router := gin.Default()

	// Register routes
	routes.RegisterRoutes(router)

	// Get server port from config
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = ":8080"
	}

	log.Printf("🚀 Server running on %s", port)
	router.Run(port)
}
`

// EnvTemplate is the .env.example template
var EnvTemplate = `APP_NAME={{.AppName}}
APP_ENV=local
APP_DEBUG=true

# Database Configuration
DB_HOST=localhost
DB_PORT=3306
DB_DATABASE={{.AppName}}
DB_USERNAME=root
DB_PASSWORD=

# Server Configuration
SERVER_PORT=:8080
`

// ReadmeTemplate is the README.md template
var ReadmeTemplate = `# {{.AppName}}

Created with **Gophant** 🐘🐹 - Scaffolder for Go web developers

## Getting Started

### Prerequisites
- Go 1.21 or higher
- MySQL (default) or PostgreSQL (optional)

### Installation

~~~bash
go mod tidy
go mod download
~~~

### Running the Application

~~~bash
cp .env.example .env
go run cmd/main.go
~~~

The server will start on the port defined in \~.env\~ (default: http://localhost:8080).

> If no \~.env\~ file is found, default configuration values will be used.

---

## Project Structure

~~~
{{.AppName}}/
├── cmd/
│   └── main.go            # Application entry point
├── internal/
│   ├── models/            # GORM models
│   ├── handlers/          # Request handlers
│   ├── routes/            # Route definitions
│   ├── middleware/        # Middleware (auth, cors, etc)
│   └── database/          # Database connections
├── pkg/
│   ├── config/            # Configuration management
│   └── validators/        # Validation logic
├── migrations/            # Database migrations
├── .env.example           # Environment variables template
└── go.mod                 # Go module definition
~~~

---

## Gophant Features

✅ Pre-configured Gin routes  
✅ Environment variable setup  
✅ Modular project structure  
✅ GORM integration ready  
✅ MySQL support out of the box  
✅ Middleware support  
✅ Clean architecture  

---


## Environment Configuration


~~~bash
cp .env.example .env
~~~

Update values based on your setup.

---

## Adding Models


~~~go
package models

type User struct {
	ID    uint
	Name  string
	Email string
}
~~~

---

## Adding Handlers


~~~go
package handlers

import "github.com/gin-gonic/gin"

func GetUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "List of users",
	})
}
~~~

---

## Adding Routes


~~~go
import "{{.AppName}}/internal/handlers"

v1.GET("/users", handlers.GetUsers)
~~~

---


---

## Resources

- https://gin-gonic.com/
- https://gorm.io/
- https://golang.org/doc/

---

## License

MIT License - {{.Year}}

---

Created with Gophant 🐘🐹
`

// ConfigTemplate is the pkg/config/config.go template
var ConfigTemplate = `package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config holds application configuration
type Config struct {
	AppName    string
	AppEnv     string
	AppDebug   bool
	DBHost     string
	DBPort     string
	DBDatabase string
	DBUsername string
	DBPassword string
	ServerPort string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	// Load .env file if it exists
	godotenv.Load()

	return &Config{
		AppName:    getEnv("APP_NAME", "gophant-app"),
		AppEnv:     getEnv("APP_ENV", "local"),
		AppDebug:   getEnv("APP_DEBUG", "true") == "true",
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBDatabase: getEnv("DB_DATABASE", "gophant"),
		DBUsername: getEnv("DB_USERNAME", "root"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		ServerPort: getEnv("SERVER_PORT", ":8080"),
	}
}

// getEnv gets an environment variable with a default value
func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
`

// RoutesTemplate is the internal/routes/routes.go template
var RoutesTemplate = `package routes

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers all application routes
func RegisterRoutes(router *gin.Engine) {
	// Health check endpoint
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"message": "Welcome to Gophant API",
		})
	})

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// Example endpoint
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "healthy",
			})
		})

		// Add your routes here
		// Example:
		// v1.GET("/users", handlers.GetUsers)
		// v1.POST("/users", handlers.CreateUser)
	}
}
`

// GitignoreTemplate is the .gitignore template
var GitignoreTemplate = `# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.so.*
*.dylib

*.test

# Output of the go coverage tool
*.out

# Go workspace file
go.work

# Dependency directories (remove the comment below to include it)
vendor/

# Environment variables
.env
.env.local
.env.*.local

# IDE
.vscode/
.idea/
*.swp
*.swo
*~
.DS_Store

# Air - Live reload for Go
.air.toml
tmp/
`
