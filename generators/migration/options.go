package migration

import (
	"github.com/swiftcarrot/dashi/generators/attrs"
	"github.com/swiftcarrot/flect"
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
