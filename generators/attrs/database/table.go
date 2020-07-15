package database

import "github.com/swiftcarrot/flect"

type Table struct {
	Columns []Column
	Name    flect.Ident
}
