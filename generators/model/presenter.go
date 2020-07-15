package model

import (
	"github.com/swiftcarrot/dashi/generators/attrs"
	"github.com/swiftcarrot/flect/name"
)

type presenter struct {
	Name        name.Ident
	Encoding    name.Ident
	Imports     []string
	Validations attrs.Attrs
}
