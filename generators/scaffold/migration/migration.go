package migration

import (
	"path/filepath"
	"strings"
	"time"

	"github.com/gobuffalo/fizz"
	"github.com/gobuffalo/genny/v2"
	"github.com/swiftcarrot/dashi/generators/scaffold"
)

func getTime() string {
	t := time.Now()

	return t.Format("20060102150405")
}

func New(opts *scaffold.Options) (*genny.Generator, error) {
	g := genny.New()
	t := fizz.NewTable(opts.Name.Underscore().Pluralize().String(), map[string]interface{}{
		"timestamps": true,
	})
	for _, attr := range opts.Attrs {
		o := fizz.Options{}
		name := attr.Name.Underscore().String()
		colType := attr.FizzType()
		if name == "id" {
			o["primary"] = true
		}
		if strings.HasPrefix(attr.GoType(), "nulls.") {
			o["null"] = true
		}
		if err := t.Column(name, colType, o); err != nil {
			return g, err
		}
	}
	var f genny.File
	up := t.Fizz()
	down := t.UnFizz()
	f = genny.NewFileS(filepath.Join("migrations", getTime()+"_"+opts.Name.Underscore().String()+".up.fizz"), up)
	g.File(f)
	f = genny.NewFileS(filepath.Join("migrations", getTime()+"_"+opts.Name.Underscore().String()+".down.fizz"), down)
	g.File(f)
	return g, nil

}
