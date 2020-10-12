package provider

import (
	"context"

	"github.com/matryer/resync"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoClient *mongo.Client
	errMongo    error
	runMongo    resync.Once
)

// MongoDB struct to create new mongo db connection client
//
// Address: the URI of mongo db (mongodb://localhost:27017)
// Username: the username for authentication.
// Password: the password for authentication.
// Database: the name of the database to use for authentication.
// MaxPoolSize: specifies that maximum number of connections allowed in the driver's connection pool to each server.
type MongoDB struct {
	Address     string
	Username    string
	Password    string
	Database    string
	MaxPoolSize int
}

// Client singleton of Mongo DB client connector, set MongoDB struct to call this method
// library with go.mongodb.org/mongo-driver/mongo
func (m MongoDB) Client() (*mongo.Client, error) {
	auth := options.Credential{
		AuthSource: m.Database,
		Username:   m.Username,
		Password:   m.Password,
	}

	runMongo.Do(func() {
		mongoClient, errMongo = mongo.Connect(context.Background(), options.Client().
			SetAuth(auth).
			SetMaxPoolSize(uint64(m.MaxPoolSize)).
			ApplyURI(m.Address))
	})

	if err := mongoClient.Ping(context.Background(), nil); err != nil {
		return nil, err
	}

	return mongoClient, errMongo
}

// Reset singleton mongo db connection client
func (m MongoDB) Reset() {
	runMongo.Reset()
}
