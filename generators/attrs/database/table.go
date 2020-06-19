package database

import "github.com/gobuffalo/flect"

type Table struct {
	Columns    []Column
	Name       flect.Ident
	PrimaryKey []string
}
