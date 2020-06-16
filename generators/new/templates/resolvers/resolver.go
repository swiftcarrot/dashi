package resolvers

import "github.com/gobuffalo/pop"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Conn *pop.Connection
}

var PerPage = 20
var Page = 1
