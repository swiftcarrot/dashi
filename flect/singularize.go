package flect

import (
	"sync"
)

var singularMoot = &sync.RWMutex{}

func singularize(str string) string {
	for _, inflection := range compiledSingularMaps {
		if inflection.regexp.MatchString(str) {
			return inflection.regexp.ReplaceAllString(str, inflection.replace)
		}
	}
	return str
}

// Singularize returns a singular version of the string
//	users = user
//	data = datum
//	people = person
func Singularize(s string) string {
	return New(s).Singularize().String()
}

// Singularize returns a singular version of the string
//	users = user
//	data = datum
//	people = person
func (i Ident) Singularize() Ident {
	s := i.Original
	if len(s) == 0 {
		return i
	}

	singularMoot.RLock()
	defer singularMoot.RUnlock()

	if p, ok := pluralToSingle[s]; ok {
		return New(p)
	}
	if _, ok := singleToPlural[s]; ok {
		return i
	}

	return New(singularize(s))
}
