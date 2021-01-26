package neo4j

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dynastymasra/cookbook/provider"

	j "github.com/neo4j/neo4j-go-driver/neo4j"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/neo4j"
	"github.com/sirupsen/logrus"
)

const (
	migrationSourcePath = "file://migration"
	migrationFilePath   = "./migration"
)

// CreateMigrationFiles for Neo4J and put in migration folder
func CreateMigrationFiles(filename string) error {
	if len(filename) == 0 {
		return errors.New("migration filename is not provided")
	}

	timeStamp := time.Now().Unix()
	upMigrationFilePath := fmt.Sprintf("%s/%d_%s.up.cypher", migrationFilePath, timeStamp, filename)
	downMigrationFilePath := fmt.Sprintf("%s/%d_%s.down.cypher", migrationFilePath, timeStamp, filename)

	if err := provider.CreateFile(upMigrationFilePath); err != nil {
		return err
	}
	logrus.Println("created", upMigrationFilePath)

	if err := provider.CreateFile(downMigrationFilePath); err != nil {
		os.Remove(upMigrationFilePath)
		return err
	}

	logrus.Println("created", downMigrationFilePath)

	return nil
}

// Migration preparation for Neo4J
func Migration(client j.Driver) (*migrate.Migrate, error) {
	config := &neo4j.Config{
		MigrationsLabel: neo4j.DefaultMigrationsLabel,
		MultiStatement:  true,
	}

	driver, err := neo4j.WithInstance(client, config)
	if err != nil {
		logrus.WithError(err).Errorln("Failed open instance")

		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(migrationSourcePath, "neo4j", driver)
	if err != nil {
		logrus.WithError(err).Errorln("Failed migration data")

		return nil, err
	}

	return m, nil
}
