package dashboard

import "github.com/swiftcarrot/dashi/flect"

// Options for generating dashboard
type Options struct {
	Name flect.Ident
}

// Validate that options are usuable
func (opts *Options) Validate() error {
	return nil
}
