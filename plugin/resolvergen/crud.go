package resolvergen

import (
	"bytes"
	_ "embed"
	"strings"
	"text/template"

	"github.com/swiftcarrot/flect"
	"github.com/swiftcarrot/gqlgen/codegen"
	"github.com/swiftcarrot/gqlgen/codegen/templates"
)

//go:embed crud.tmpl
var crudTemplate string

type crudResolver struct {
	Field *codegen.Field
}

func renderCRUD(field *codegen.Field) (string, error) {
	t, err := template.New("crud.tmpl").Funcs(template.FuncMap{
		"lcFirst":          templates.LcFirst,
		"go":               templates.ToGo,
		"isCRUDResolver":   isCRUDResolver,
		"entityFromResult": func(s string) string { return strings.ReplaceAll(s, "Items", "") },
		"pluralize":        func(s string) string { return flect.Pluralize(s) },
	}).Parse(crudTemplate)
	if err != nil {
		return "", err
	}

	var code bytes.Buffer

	err = t.Execute(&code, &crudResolver{
		Field: field,
	})
	if err != nil {
		return "", err
	}

	return strings.Trim(code.String(), "\t \n"), nil
}

func isCRUDResolver(field codegen.Field, kind string) bool {
	for _, direct := range field.Directives {
		if direct.Name == "generated" && direct.Args[0].Value == kind {
			return true
		}
	}
	return false
}
