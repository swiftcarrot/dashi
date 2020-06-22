package database

import (
	"github.com/gobuffalo/flect"
	"github.com/gobuffalo/nulls"
)

//TODO implement default and unique
type Column struct {
	Name           flect.Ident
	ColType        string
	Nullable       bool
	Default        nulls.String
	Unique         bool
	Primary        bool
	IsSequence     bool
	SequenceSuffix string
}
