package attrs

import (
	"errors"
	"strings"

	"github.com/swiftcarrot/dashi/flect"
)

// ErrRepeatedAttr is returned when parsing an array with repeated names
var ErrRepeatedAttr = errors.New("duplicate attr name")

// Parse takes a string like name:inputType:goType and turns it into an Attr
func Parse(arg string) (Attr, error) {
	arg = strings.TrimSpace(arg)
	attr := Attr{
		Original:  arg,
		inputType: "string",
	}
	if len(arg) == 0 {
		return attr, errors.New("argument can not be blank")
	}

	parts := strings.Split(arg, ":")
	attr.Name = flect.New(parts[0])
	if len(parts) > 1 {
		attr.inputType = parts[1]
	}

	if len(parts) > 2 {
		attr.goType = parts[2]
	}

	return attr, nil
}

func ParseArgs(args ...string) (Attrs, error) {
	var attrs Attrs
	for _, arg := range args {
		var attr, _ = Parse(arg)
		attrs = append(attrs, attr)
	}
	return attrs, nil
}
