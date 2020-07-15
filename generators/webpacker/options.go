package webpacker

import "github.com/swiftcarrot/flect"

// Options for generating webpacker
type Options struct {
	Name flect.Ident
}

// Validate that options are usuable
func (opts *Options) Validate() error {
	return nil
}
