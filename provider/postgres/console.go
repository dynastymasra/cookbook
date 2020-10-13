package postgres

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	migrationSourcePath = "file://migration"
	migrationFilePath   = "./migration"
)

func CreateMigrationFiles(filename string) error {
	if len(filename) == 0 {
		return errors.New("migration filename is not provided")
	}

	timestamp := time.Now().Unix()
	upMigrationFilePath := fmt.Sprintf("%s/%d_%s.up.sql", migrationFilePath, timestamp, filename)
	downMigrationFilePath := fmt.Sprintf("%s/%d_%s.down.sql", migrationFilePath, timestamp, filename)

	if err := createFile(upMigrationFilePath); err != nil {
		return err
	}
	log.Println("created", upMigrationFilePath)

	if err := createFile(downMigrationFilePath); err != nil {
		os.Remove(upMigrationFilePath)
		return err
	}

	log.Println("created", downMigrationFilePath)

	return nil
}

func Migration(data *gorm.DB) (*migrate.Migrate, error) {
	db, err := data.DB()
	if err != nil {
		return nil, err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		logrus.WithError(err).Errorln("Failed open instance")

		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(migrationSourcePath, "postgres", driver)
	if err != nil {
		logrus.WithError(err).Errorln("Failed migration data")

		return nil, err
	}

	return m, nil
}

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

func createFile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}

	return nil
}
