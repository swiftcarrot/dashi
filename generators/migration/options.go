package migration

// Options for generating migration
type Options struct {
	Name string
	Time string
}

// Validate that options are usuable
func (opts *Options) Validate() error {
	return nil
}
