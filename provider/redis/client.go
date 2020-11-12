package redis

import (
	"context"

	"github.com/matryer/resync"

	"github.com/go-redis/redis/v8"
)

var (
	client *redis.Client
	err    error
	once   resync.Once
)

// Config struct to create new redis connection client
//
// Address: the redis address
// Password: the redis password
// Database: redis database default value is 0
// PoolSize: Maximum number of socket connections
// MinIdleConn: Minimum number of idle connections which is useful when establishing
type Config struct {
	Address     string
	Password    string
	Database    int
	PoolSize    int
	MinIdleConn int
}

// Client create new redis client connection
func (r Config) Client() (*redis.Client, error) {
	once.Do(func() {
		client = redis.NewClient(&redis.Options{
			Addr:         r.Address,
			Password:     r.Password,
			DB:           r.Database,
			PoolSize:     r.PoolSize,
			MinIdleConns: r.MinIdleConn,
		})

		err = client.Ping(context.Background()).Err()
	})

	if err := r.Ping(); err != nil {
		return nil, err
	}

	return client, err
}

// Ping check database connection status
func (r Config) Ping() error {
	return client.Ping(context.Background()).Err()
}

// Reset reset redis client connection
func (r Config) Reset() {
	once.Reset()
}
