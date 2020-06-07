package new

// Options for generating new
type Options struct {
	Name    string
	Package string
}

// Validate that options are usuable
func (opts *Options) Validate() error {
	return nil
}
