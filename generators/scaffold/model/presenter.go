package model

import (
	"github.com/gobuffalo/flect/name"
	"github.com/swiftcarrot/dashi/generators/scaffold/attrs"
)

type presenter struct {
	Name        name.Ident
	Encoding    name.Ident
	Imports     []string
	Validations attrs.Attrs
}
