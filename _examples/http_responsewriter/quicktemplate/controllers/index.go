package controllers

import (
	"github.com/nooberfsh/iris/_examples/http_responsewriter/quicktemplate/templates"

	"github.com/nooberfsh/iris"
)

// Index renders our ../templates/index.qtpl file using the compiled ../templates/index.qtpl.go file.
func Index(ctx iris.Context) {
	tmpl := &templates.Index{}

	// render the template
	ExecuteTemplate(ctx, tmpl)
}
