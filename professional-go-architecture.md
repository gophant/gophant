# Professional Go Application Architecture: Go + Laravel Best Practices

## The Complete, Professional Structure

```
myapp/
├── cmd/                          ← CLI & Application Entry Points (GO STANDARD)
│   ├── server/
│   │   └── main.go              ← Start HTTP server
│   ├── cli/
│   │   └── main.go              ← CLI tool entry point
│   └── migrate/
│       └── main.go              ← Database migration runner
│
├── pkg/                          ← Reusable Packages (GO STANDARD)
│   ├── config/
│   │   └── config.go            ← Configuration management
│   ├── logger/
│   │   └── logger.go            ← Logging (Zap, Logrus, etc.)
│   ├── database/
│   │   ├── mysql.go             ← Database connection
│   │   └── migration.go         ← Migration runner
│   ├── cache/
│   │   ├── redis.go             ← Redis client
│   │   └── cache.go             ← Caching interface
│   ├── middleware/
│   │   ├── auth.go              ← Auth middleware
│   │   ├── cors.go              ← CORS middleware
│   │   └── logging.go           ← Logging middleware
│   └── utils/
│       ├── validation.go        ← Input validation
│       └── response.go          ← API response formatting
│
├── internal/                     ← Private Application Code (GO STANDARD)
│   ├── models/
│   │   ├── user.go              ← User model
│   │   └── post.go              ← Post model
│   ├── handlers/                ← (Same as Laravel Controllers)
│   │   ├── user.go              ← User handlers (GET, POST, PUT, DELETE)
│   │   └── post.go              ← Post handlers
│   ├── services/                ← Business Logic (Laravel Service Providers)
│   │   ├── user_service.go      ← User business logic
│   │   └── post_service.go      ← Post business logic
│   ├── repositories/            ← Data Access Layer
│   │   ├── user_repository.go   ← User data queries
│   │   └── post_repository.go   ← Post data queries
│   └── cli/                      ← CLI Commands (Cobra)
│       ├── commands/
│       │   ├── root.go          ← Root command
│       │   ├── model.go         ← go run cmd/cli/main.go model create User
│       │   ├── migration.go     ← go run cmd/cli/main.go migration create
│       │   ├── seeder.go        ← go run cmd/cli/main.go seed run
│       │   └── tinker.go        ← go run cmd/cli/main.go tinker (REPL)
│       └── generators/
│           ├── model.go         ← Generate model file
│           ├── migration.go     ← Generate migration file
│           └── seeder.go        ← Generate seeder file
│
├── app/                          ← Laravel-Style Organization (OPTIONAL, FOR CLARITY)
│   ├── models/          → Symlink to internal/models
│   ├── handlers/        → Symlink to internal/handlers
│   ├── services/        → Symlink to internal/services
│   └── middleware/      → Symlink to pkg/middleware
│
├── config/                       ← Configuration Files (LIKE LARAVEL)
│   ├── app.go           ← Application configuration
│   ├── database.go      ← Database configuration
│   ├── cache.go         ← Redis/Cache configuration
│   └── logging.go       ← Logging configuration
│
├── database/                     ← Database Files (LIKE LARAVEL)
│   ├── migrations/
│   │   ├── 001_create_users_table.up.sql
│   │   ├── 001_create_users_table.down.sql
│   │   ├── 002_create_posts_table.up.sql
│   │   └── 002_create_posts_table.down.sql
│   └── seeders/
│       ├── user_seeder.go       ← Generate and run with Cobra
│       └── post_seeder.go
│
├── routes/                       ← Route Definitions (LIKE LARAVEL)
│   ├── routes.go        ← All route definitions
│   └── api.go           ← API-specific routes (future)
│
├── resources/                    ← Frontend Assets (LIKE LARAVEL)
│   ├── views/           ← HTML templates (if needed)
│   ├── lang/            ← Language files
│   └── css/, js/        ← Static assets
│
├── storage/                      ← File Storage (LIKE LARAVEL)
│   ├── logs/            ← Log files
│   ├── cache/           ← Cache files
│   └── uploads/         ← User uploads
│
├── tests/                        ← Test Files (GO STANDARD)
│   ├── unit/            ← Unit tests
│   ├── feature/         ← Feature tests
│   └── integration/     ← Integration tests
│
├── .env.example         ← Environment template
├── go.mod              ← Go module definition
├── go.sum              ← Dependency lock file
├── main.go             ← (Optional, for convenience)
└── README.md
```

---

## How It All Works Together

### GO STANDARD (`cmd/`, `pkg/`, `internal/`)

**Why this matters:**
- ✅ Go developers expect this structure
- ✅ Clear separation of concerns
- ✅ Industry standard
- ✅ Follows Go conventions

### LARAVEL FAMILIARITY (`app/`, `config/`, `database/`, `routes/`)

**Why this matters:**
- ✅ Laravel developers feel at home
- ✅ Clear organization
- ✅ Predictable file locations
- ✅ Easy onboarding

---

## SYNCING cmd WITH PROJECT STRUCTURE

### The Key: cmd/ runs everything, pkg/ & internal/ do the work

```
cmd/server/main.go
    ↓
imports pkg/ & internal/
    ↓
├── pkg/config/      ← Load config
├── pkg/logger/      ← Initialize logger
├── pkg/database/    ← Connect to DB
├── pkg/cache/       ← Connect to Redis
├── internal/handlers/  ← Handle routes
└── routes/routes.go    ← Define routes
```

---

## Example 1: Start Server

```bash
go run cmd/server/main.go
```

**What happens:**
```go
// cmd/server/main.go
package main

import (
    "myapp/pkg/config"
    "myapp/pkg/logger"
    "myapp/pkg/database"
    "myapp/pkg/cache"
    "myapp/routes"
    "myapp/internal/handlers"
)

func main() {
    // Load config
    cfg := config.Load()
    
    // Initialize logger (GO EXCELLENCE)
    log := logger.New(cfg.LogLevel)
    
    // Connect to database
    db := database.Connect(cfg.DatabaseURL, log)
    
    // Connect to Redis (GO EXCELLENCE)
    cache := cache.NewRedis(cfg.RedisURL, log)
    
    // Setup handlers with dependencies
    h := handlers.New(db, cache, log)
    
    // Define routes
    router := routes.Setup(h)
    
    // Start server
    log.Info("Server started on :8080")
    router.Run(":8080")
}
```

---

## Example 2: Run Migrations

```bash
go run cmd/migrate/main.go up
```

**What happens:**
```go
// cmd/migrate/main.go
package main

import (
    "myapp/pkg/database"
    "myapp/pkg/logger"
    "flag"
)

func main() {
    direction := flag.String("direction", "up", "up or down")
    
    cfg := config.Load()
    log := logger.New(cfg.LogLevel)
    db := database.Connect(cfg.DatabaseURL, log)
    
    migrator := database.NewMigrator(db, log)
    migrator.Run(*direction)
}
```

---

## Example 3: Run CLI Commands

```bash
go run cmd/cli/main.go model create User name:string email:string
go run cmd/cli/main.go migration create create_users_table
go run cmd/cli/main.go seed run
```

**What happens:**
```go
// cmd/cli/main.go
package main

import (
    "myapp/internal/cli/commands"
    "myapp/pkg/database"
    "myapp/pkg/logger"
)

func main() {
    cfg := config.Load()
    log := logger.New(cfg.LogLevel)
    db := database.Connect(cfg.DatabaseURL, log)
    
    // Setup CLI with dependencies
    cli := commands.New(db, log)
    cli.Execute()
}
```

---

## HOW pkg/ IS UTILIZED

### pkg/ Contains Reusable, Shareable Code

```
pkg/config/     ← Configuration loading (reusable)
pkg/logger/     ← Logging (reusable)
pkg/database/   ← DB connection (reusable)
pkg/cache/      ← Redis client (reusable)
pkg/middleware/ ← HTTP middleware (reusable)
pkg/utils/      ← Utilities (reusable)
```

**Why this matters:**
- ✅ Can be exported to other Go projects
- ✅ Can be published as packages
- ✅ Clean, isolated responsibilities
- ✅ Easy to test

### Example: Using pkg/logger Everywhere

```go
// cmd/server/main.go
log := logger.New(cfg.LogLevel)

// pkg/database/mysql.go
func Connect(url string, log *logger.Logger) {
    log.Info("Connecting to database...")
}

// internal/handlers/user.go
func (h *Handler) GetUser(c *gin.Context) {
    h.log.Debug("Getting user...")
}

// pkg/cache/redis.go
func NewRedis(url string, log *logger.Logger) {
    log.Info("Connecting to Redis...")
}
```

**Same logger everywhere! Consistent logging!**

---

## GO EXCELLENCE: Logging + Redis

### Logging (Go has EXCELLENT logging libraries)

```go
// pkg/logger/logger.go
import "go.uber.org/zap"

type Logger struct {
    *zap.Logger
}

func New(level string) *Logger {
    cfg := zap.NewProductionConfig()
    cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
    
    logger, _ := cfg.Build()
    return &Logger{logger}
}

// Usage
log.Info("user created", zap.String("email", "user@example.com"))
log.Error("failed to save", zap.Error(err))
```

**Why Go's logging is excellent:**
- ✅ Structured logging (JSON output)
- ✅ Levels (Debug, Info, Warn, Error)
- ✅ Performance optimized
- ✅ Production-ready out of the box

**Laravel doesn't have this level of built-in excellence!**

---

### Redis Caching (Go handles Redis beautifully)

```go
// pkg/cache/redis.go
import "github.com/go-redis/redis/v8"

type Cache struct {
    client *redis.Client
}

func NewRedis(url string, log *logger.Logger) *Cache {
    client := redis.NewClient(&redis.Options{
        Addr: url,
    })
    
    log.Info("Connected to Redis")
    return &Cache{client}
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) error {
    return c.client.Set(context.Background(), key, value, ttl).Err()
}

func (c *Cache) Get(key string) (interface{}, error) {
    return c.client.Get(context.Background(), key).Result()
}
```

**Usage:**
```go
// Caching user data
cache.Set("user:1", user, time.Hour)
user, _ := cache.Get("user:1")
```

**Why Go handles Redis better:**
- ✅ Native connection pooling
- ✅ Async support (goroutines)
- ✅ Better performance
- ✅ Simpler API

---

## SYNCING cmd WITH app/

### Option 1: Direct Import (Simplest)

```go
// cmd/server/main.go
import "myapp/internal/handlers"

h := handlers.New(db, cache, log)
```

---

### Option 2: Symbolic Links (If You Want App Directory)

```bash
# Create symlinks
ln -s ../internal/models app/models
ln -s ../internal/handlers app/handlers
ln -s ../internal/services app/services
ln -s ../pkg/middleware app/middleware
```

Then import either way:
```go
import (
    "myapp/app/models"      // Via symlink
    // OR
    "myapp/internal/models" // Direct
)
```

**Both work! Use what you prefer.**

---

## THE COMPLETE WORKFLOW

### Step 1: Initialize Project
```bash
go mod init github.com/yourname/myapp

# Create structure
mkdir -p cmd/{server,cli,migrate} pkg/{config,logger,database,cache,middleware} internal/{models,handlers,services,repositories,cli/{commands,generators}}
```

---

### Step 2: Configure Everything

```go
// config/app.go
type Config struct {
    AppName string
    LogLevel string
    ServerPort string
    DatabaseURL string
    RedisURL string
}

func Load() *Config {
    return &Config{
        AppName: os.Getenv("APP_NAME"),
        LogLevel: os.Getenv("LOG_LEVEL"),
        // ...
    }
}
```

---

### Step 3: Setup Dependencies

```go
// pkg/database/mysql.go
// pkg/logger/logger.go
// pkg/cache/redis.go
// All initialization logic
```

---

### Step 4: Write Handlers

```go
// internal/handlers/user.go
type Handler struct {
    db *sql.DB
    cache *cache.Cache
    log *logger.Logger
}

func (h *Handler) GetUsers(c *gin.Context) {
    // Handle request
}
```

---

### Step 5: Define Routes

```go
// routes/routes.go
func Setup(h *handlers.Handler) *gin.Engine {
    router := gin.Default()
    
    api := router.Group("/api")
    {
        api.GET("/users", h.GetUsers)
        api.POST("/users", h.CreateUser)
    }
    
    return router
}
```

---

### Step 6: Run Server

```bash
go run cmd/server/main.go
```

---

### Step 7: Use CLI Tools

```bash
# Create model
go run cmd/cli/main.go model create User

# Run migration
go run cmd/migrate/main.go up

# Seed database
go run cmd/cli/main.go seed run
```

---

## HOW THIS BECOMES GLOBALLY ACCEPTED

### Why This Architecture Would Be Accepted

1. **Combines Best of Both Worlds**
   - ✅ Go standard structure (cmd/, pkg/, internal/)
   - ✅ Laravel familiarity (app/, config/, database/)
   - ✅ Clear organization

2. **Showcases Go Excellence**
   - ✅ Logging (Zap, Logrus)
   - ✅ Caching (Redis)
   - ✅ Concurrency (goroutines)
   - ✅ Performance

3. **Developer Friendly**
   - ✅ Laravel devs understand it
   - ✅ Go devs respect it
   - ✅ Easy to learn
   - ✅ Predictable

4. **Production Ready**
   - ✅ Separation of concerns
   - ✅ Dependency injection
   - ✅ Testable code
   - ✅ Scalable

---

## SUMMARY: The Perfect Hybrid Architecture

```
GO STANDARD (cmd/, pkg/, internal/)
    +
LARAVEL FAMILIARITY (app/, config/, database/, routes/)
    +
GO EXCELLENCE (Logging, Redis, Concurrency)
    =
GLOBALLY ACCEPTED STANDARD
```

This architecture would:
- ✅ Attract Laravel developers to Go
- ✅ Make Go developers proud
- ✅ Be production-ready
- ✅ Become a new standard
- ✅ Help teams onboard faster
- ✅ Reduce friction between PHP and Go teams

---

## THIS IS THE GOPHANT VISION

When you run:
```bash
gophant create myapp
```

You should generate this entire structure with:
- ✅ Proper config loading
- ✅ Logging setup
- ✅ Database connection
- ✅ Redis setup
- ✅ CLI commands ready
- ✅ Migration system ready
- ✅ Seeder system ready
- ✅ Example handlers
- ✅ Example routes
- ✅ Ready to develop!

**Users just add their business logic. Everything else is done.**

This is what Laravel does. This is what Gophant should do.

**This is how you change the world of Go.** 🚀
