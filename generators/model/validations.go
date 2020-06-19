package model

import "github.com/swiftcarrot/dashi/generators/attrs"

func validatable(ats attrs.Attrs) attrs.Attrs {
	var xats attrs.Attrs
	for _, a := range ats {
		n := a.Name.Singularize().Pascalize().String()
		if n == "CreatedAt" || n == "UpdatedAt" {
			continue
		}
		switch a.GoType() {
		case "string", "time.Time", "int":
			xats = append(xats, a)
		}
	}
	return xats
}
