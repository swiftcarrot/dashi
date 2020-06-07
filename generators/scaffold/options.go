package scaffold

import (
	"github.com/gobuffalo/flect"
	"github.com/swiftcarrot/dashi/generators/scaffold/attrs"
)

// Options for generating scaffold
type Options struct {
	Name  flect.Ident
	Attrs attrs.Attrs
	// add your stuff here
}

// Validate that options are usuable
func (opts *Options) Validate() error {
	return nil
}
