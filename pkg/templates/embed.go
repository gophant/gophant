package templates

import _ "embed"

// Embedded templates (default)

//go:embed default/go.mod.tmpl
var goMod string

//go:embed default/main.go.tmpl
var mainGo string

//go:embed default/.env.example.tmpl
var envExample string

//go:embed default/README.md.tmpl
var readme string

//go:embed default/pkg_config_config.go.tmpl
var configGo string

//go:embed default/app_routes_routes.go.tmpl
var routesGo string

//go:embed default/cmd_root_root.go.tmpl
var rootGo string

//go:embed default/cmd_serve_serve.go.tmpl
var serveGo string

//go:embed default/gitignore.tmpl
var gitignore string

// Embedded template set exposed to loader
var Embedded = TemplateSet{
	GoMod:     goMod,
	Main:      mainGo,
	Env:       envExample,
	Readme:    readme,
	Config:    configGo,
	Routes:    routesGo,
	Root:      rootGo,
	Serve:     serveGo,
	Gitignore: gitignore,
}

// TemplateSet holds templates
type TemplateSet struct {
	GoMod     string
	Main      string
	Env       string
	Readme    string
	Config    string
	Routes    string
	Root      string
	Serve     string
	Gitignore string
}
