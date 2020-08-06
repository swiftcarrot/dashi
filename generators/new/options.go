package new

import "github.com/swiftcarrot/flect"

// Options for generating new
type Options struct {
	Name    flect.Ident
	Package string
	APIOnly bool
}

// Validate that options are usuable
func (opts *Options) Validate() error {
	return nil
}
