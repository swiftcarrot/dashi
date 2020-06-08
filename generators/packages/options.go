package packages

// Options for generating packages
type Options struct {
	Name string
}

// Validate that options are usuable
func (opts *Options) Validate() error {
	return nil
}
