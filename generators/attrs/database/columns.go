package database

import (
	"github.com/gobuffalo/nulls"
	"github.com/swiftcarrot/dashi/flect"
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
