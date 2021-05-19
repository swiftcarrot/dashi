package database

import (
	"bytes"
	"embed"

	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
)

type EmbedMigrator struct {
	pop.Migrator
	fs *embed.FS
}

func NewEmbedMigrator(fs *embed.FS, c *pop.Connection) (EmbedMigrator, error) {
	migrator := EmbedMigrator{
		Migrator: pop.NewMigrator(c),
		fs:       fs,
	}

	runner := func(mf pop.Migration, tx *pop.Connection) error {
		b, err := migrator.fs.ReadFile(mf.Path)
		if err != nil {
			return err
		}

		content, err := pop.MigrationContent(mf, tx, bytes.NewReader(b), true)
		if err != nil {
			return errors.Wrapf(err, "error processing %s", mf.Path)
		}
		if content == "" {
			return nil
		}

		err = tx.RawQuery(content).Exec()
		if err != nil {
			return errors.Wrapf(err, "error executing %s, sql: %s", mf.Path, content)
		}
		return nil
	}

	err := migrator.findMigrations(runner)
	if err != nil {
		return migrator, err
	}

	return migrator, nil
}

func (migrator *EmbedMigrator) findMigrations(runner func(mf pop.Migration, tx *pop.Connection) error) error {
	files, err := migrator.fs.ReadDir(("migrations"))
	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() {
			match, err := pop.ParseMigrationFilename(file.Name())
			if err != nil {
				return err
			}
			if match == nil {
				return nil
			}

			mf := pop.Migration{
				Path:      "migrations/" + file.Name(),
				Version:   match.Version,
				Name:      match.Name,
				DBType:    match.DBType,
				Direction: match.Direction,
				Type:      match.Type,
				Runner:    runner,
			}
			migrator.Migrations[mf.Direction] = append(migrator.Migrations[mf.Direction], mf)
		}
	}

	return nil
}
