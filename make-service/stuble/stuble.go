package stuble

import _ "embed"

//go:embed cmd.tpl
var Cmd string

//go:embed dto.tpl
var Dto string

//go:embed model.tpl
var Model string

//go:embed  service.tpl
var Service string

//go:embed handler.tpl
var Handler string

//go:embed http.tpl
var Http string

//go:embed gitignore.tpl
var Gitignore string

//go:embed env.tpl
var Env string

//go:embed README.tpl
var README string

//go:embed tepmlate.go.tmpl
var GoTest string

func GetTemplate() string {
	return GoTest
}
