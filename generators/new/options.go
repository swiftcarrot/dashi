package new

import "github.com/gobuffalo/flect"

// Options for generating new
type Options struct {
	Name    flect.Ident
	Package string
}

// Validate that options are usuable
func (opts *Options) Validate() error {
	return nil
}
