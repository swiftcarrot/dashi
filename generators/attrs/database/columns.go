package database

import "github.com/gobuffalo/flect"

//TODO implement default and unique
type Column struct {
	Name     flect.Ident
	ColType  string
	Nullable bool
	Default  string
	Unique   bool
}
