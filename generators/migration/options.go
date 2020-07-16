package migration

import (
	"github.com/swiftcarrot/dashi/flect"
	"github.com/swiftcarrot/dashi/generators/attrs"
)

// Options for generating migration
type Options struct {
	Dialect string
	Name    flect.Ident
	Time    string
	Attrs   attrs.Attrs `json:"attrs"`
}

// Validate that options are usuable
func (opts *Options) Validate() error {
	return nil
}
