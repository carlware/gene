package funcs

import (
	"strings"
	"text/template"
)

var Funcs = template.FuncMap{
	"title": strings.Title,
}
