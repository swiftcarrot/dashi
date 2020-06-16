package scaffold

import (
	"github.com/gobuffalo/flect"
	"github.com/swiftcarrot/dashi/generators/attrs"
)

// Options for generating scaffold
type Options struct {
	Name  flect.Ident
	Attrs attrs.Attrs
}

// Validate that options are usuable
func (opts *Options) Validate() error {
	return nil
}
