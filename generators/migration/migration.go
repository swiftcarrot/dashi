package migration

import (
	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/packr/v2"
	"github.com/swiftcarrot/dashi/generators/attrs/database"
	"github.com/swiftcarrot/dashi/genny"
	"github.com/swiftcarrot/dashi/genny/gogen"
)

type MigrationError struct {
	Message string
}

func (e *MigrationError) Error() string {
	return e.Message
}

func New(opts *Options) (*genny.Generator, error) {
	g := genny.New()

	if err := opts.Validate(); err != nil {
		return g, err
	}
	if opts.Dialect == "postgres" {
		if err := g.Box(packr.New("dashi:generators:migration:postgres", "../migration/templates/postgres")); err != nil {
			return g, err
		}
	} else {
		println("dialect migration not implemented")
		return nil, &MigrationError{Message: "dialect migration not implemented"}
	}
	var ctx map[string]interface{}
	help := map[string]interface{}{}

	if len(opts.Attrs) > 0 {
		var cols []database.Column
		var sequences []string
		var tableName = opts.Name.Pluralize().Underscore()
		for _, attr := range opts.Attrs {
			col := attr.PostgresColumn()
			if col.IsSequence {
				seqName := tableName.String() + col.SequenceSuffix
				col.Default = nulls.String{
					Valid:  true,
					String: "nextval('" + seqName + "'::regclass)",
				}
				sequences = append(sequences)
			}
			cols = append(cols, col)

		}
		table := database.Table{
			Columns: cols,
			Name:    tableName,
		}
		ctx = map[string]interface{}{
			"opts":     opts,
			"sequence": sequences,
			"table":    table,
		}
	} else {
		ctx = map[string]interface{}{}
	}

	t := gogen.TemplateTransformer(ctx, help)
	g.Transformer(t)
	g.Transformer(genny.Replace("-time-", opts.Time))

	if len(opts.Attrs) > 0 {
		g.Transformer(genny.Replace("-name-", "create_"+opts.Name.Underscore().Pluralize().String()))
	} else {
		g.Transformer(genny.Replace("-name-", opts.Name.Underscore().String()))
	}

	return g, nil
}
