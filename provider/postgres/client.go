package postgres

import (
	"database/sql"
	"fmt"

	"github.com/matryer/resync"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db   *gorm.DB
	err  error
	once resync.Once
)

// Config struct to create new postgres connection client
//
//{
//	Database: the postgres database name
//	Host: the postgres database host (localhost)
//	Port: the postgres database port (5432)
//	Username: the postgres database username
//	Password: the postgres database password
//	Params: the postgres database params, use space to separate value (sslmode=disable TimeZone=Asia/Jakarta)
//	MaxIdleConn: sets the maximum number of connections in the idle connection pool.
//	MaxOpenConn: sets the maximum number of open connections to the database.
//	LogMode: sets log mode, 1(Silent) - 2(Error) - 3(Warn) - 4(Info), default is Error
//	DebugEnabled: sets true if enabled debug mode, will show query on console
//}
type Config struct {
	Database     string
	Host         string
	Port         int
	Username     string
	Password     string
	Params       string
	MaxIdleConn  int
	MaxOpenConn  int
	LogMode      int
	DebugEnabled bool
}

// Client singleton of Postgres connection client, use Postgres struct to call this method
// library with github.com/jinzhu/gorm
func (p *Config) Client() (*gorm.DB, error) {
	once.Do(func() {
		dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d %s",
			p.Username, p.Password, p.Database, p.Host, p.Port, p.Params)

		logMode := func() logger.LogLevel {
			switch p.LogMode {
			case 1:
				return logger.Silent
			case 2:
				return logger.Error
			case 3:
				return logger.Warn
			case 4:
				return logger.Info
			default:
				return logger.Error
			}
		}

		config := &gorm.Config{
			Logger: logger.Default.LogMode(logMode()),
		}

		db, err = gorm.Open(postgres.Open(dsn), config)
		if err != nil {
			return
		}

		if p.DebugEnabled {
			db = db.Debug()
		}

		sqlDB, errPostgres := db.DB()
		if errPostgres != nil {
			return
		}

		sqlDB.SetMaxIdleConns(p.MaxIdleConn)
		sqlDB.SetMaxOpenConns(p.MaxOpenConn)

		errPostgres = sqlDB.Ping()
	})

	if err := p.Ping(); err != nil {
		return nil, err
	}

	return db, err
}

// Ping check database connection
func (p *Config) Ping() error {
	conn, err := db.DB()
	if err != nil {
		return err
	}

	return conn.Ping()
}

// Close database connection
func (p *Config) Close() error {
	conn, err := db.DB()
	if err != nil {
		return err
	}
	return conn.Close()
}

// Reset singleton postgres connection client
func (p *Config) Reset() {
	once.Reset()
}

// SetDB with existing connection
func (p *Config) SetDB(sqlDB *sql.DB) error {
	db, err = gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{})
	return err
}
