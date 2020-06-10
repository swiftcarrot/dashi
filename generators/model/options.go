package model

import (
	"github.com/gobuffalo/flect"
	"github.com/swiftcarrot/dashi/generators/scaffold/attrs"
)

// Options for generating model
type Options struct {
	Name  flect.Ident
	Attrs attrs.Attrs
}

// Validate that options are usuable
func (opts *Options) Validate() error {
	return nil
}
