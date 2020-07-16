package database

import "github.com/swiftcarrot/dashi/flect"

type Table struct {
	Columns []Column
	Name    flect.Ident
}
