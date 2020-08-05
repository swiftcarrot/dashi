package flect

import (
	"strings"
	"sync"
)

var pluralMoot = &sync.RWMutex{}

func pluralize(str string) string {
	for _, inflection := range compiledPluralMaps {
		if inflection.regexp.MatchString(str) {
			return inflection.regexp.ReplaceAllString(str, inflection.replace)
		}
	}
	return str
}

// Pluralize returns a plural version of the string
//	user = users
//	person = people
//	datum = data
func Pluralize(s string) string {
	return New(s).Pluralize().String()
}

// Pluralize returns a plural version of the string
//	user = users
//	person = people
//	datum = data
func (i Ident) Pluralize() Ident {
	s := i.Original
	if len(s) == 0 {
		return New("")
	}

	pluralMoot.RLock()
	defer pluralMoot.RUnlock()

	ls := strings.ToLower(s)
	if _, ok := pluralToSingle[ls]; ok {
		return i
	}
	if p, ok := singleToPlural[ls]; ok {
		return New(p)
	}

	return New(pluralize(s))
}
