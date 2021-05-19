package plushgen

import (
	"github.com/gobuffalo/plush/v4"
	"github.com/swiftcarrot/dashi/genny"
)

// Transformer will plushify any file that has a ".plush" extension
func Transformer(ctx *plush.Context) genny.Transformer {
	t := genny.NewTransformer(".plush", func(f genny.File) (genny.File, error) {
		s, err := plush.RenderR(f, ctx)
		if err != nil {
			return f, err
		}
		return genny.NewFileS(f.Name(), s), nil
	})
	t.StripExt = true
	return t
}
