package provider

import (
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/sirupsen/logrus"
)

func RunMigration(migration *migrate.Migrate) error {
	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		logrus.WithError(err).Errorln("Failed run database migration")
		return err
	}
	return nil
}

func RollbackMigration(migration *migrate.Migrate) error {
	if err := migration.Steps(-1); err != nil {
		logrus.WithError(err).Errorln("Failed rollback database migration")
		return err
	}
	return nil
}

func CreateFile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}

	return nil
}
