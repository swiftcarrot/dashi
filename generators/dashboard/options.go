package dashboard

// Options for generating dashboard
type Options struct {
	Name string
}

// Validate that options are usuable
func (opts *Options) Validate() error {
	return nil
}
