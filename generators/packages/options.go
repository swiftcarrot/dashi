package packages

import "github.com/swiftcarrot/dashi/flect"

// Options for generating packages
type Options struct {
	Name flect.Ident
}

// Validate that options are usuable
func (opts *Options) Validate() error {
	return nil
}
