package mongo

import (
	"context"

	"github.com/matryer/resync"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	client *mongo.Client
	err    error
	once   resync.Once
)

// Config struct to create new mongo db connection client
//
//{
//	Address: the URI of mongo db (mongodb://localhost:27017)
//	Username: the username for authentication.
//	Password: the password for authentication.
//	Database: the name of the database to use for authentication.
//	MaxPoolSize: specifies that maximum number of connections allowed in the driver's connection pool to each server.
//}
type Config struct {
	Address     string
	Username    string
	Password    string
	Database    string
	MaxPoolSize int
}

// Client singleton of Mongo DB client connector, set MongoDB struct to call this method
// library with go.mongodb.org/mongo-driver/mongo
func (m Config) Client() (*mongo.Client, error) {
	once.Do(func() {
		auth := options.Credential{
			AuthSource: m.Database,
			Username:   m.Username,
			Password:   m.Password,
		}

		client, err = mongo.Connect(context.Background(), options.Client().
			SetAuth(auth).
			SetMaxPoolSize(uint64(m.MaxPoolSize)).
			ApplyURI(m.Address))
	})

	if err := m.Ping(); err != nil {
		return nil, err
	}

	return client, err
}

// Ping mongodb connection to check database status
func (m Config) Ping() error {
	return client.Ping(context.Background(), nil)
}

// Reset singleton mongo db connection client
func (m Config) Reset() {
	once.Reset()
}
