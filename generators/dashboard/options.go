package dashboard

import "github.com/gobuffalo/flect"

// Options for generating dashboard
type Options struct {
	Name flect.Ident
}

// Validate that options are usuable
func (opts *Options) Validate() error {
	return nil
}
