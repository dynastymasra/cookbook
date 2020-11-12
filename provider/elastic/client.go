package elastic

import (
	"errors"
	"net/http"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/matryer/resync"
)

var (
	client *elasticsearch.Client
	errs   error
	once   resync.Once
)

// Config struct to create new elasticsearch connection client
//
// Address: the elasticsearch database name
// Username: the elasticsearch database username
// Password: the elasticsearch database password
// MaxIdlePerHost: sets the maximum number of idle connections to the database.
// MaxConnPerHost: sets the maximum number of open connections to the database.
type Config struct {
	Address        string
	Username       string
	Password       string
	MaxConnPerHost int
	MaxIdlePerHost int
}

// Client generate new elasticsearch client connection
func (e Config) Client() (*elasticsearch.Client, error) {
	once.Do(func() {
		addresses := strings.Split(e.Address, ",")

		config := elasticsearch.Config{
			Addresses: addresses,
			Username:  e.Username,
			Password:  e.Password,
			Transport: &http.Transport{
				MaxConnsPerHost:     e.MaxConnPerHost,
				MaxIdleConnsPerHost: e.MaxIdlePerHost,
			},
		}

		client, errs = elasticsearch.NewClient(config)
	})

	if err := e.Ping(); err != nil {
		return nil, err
	}

	return client, errs
}

// Ping check database connection status
func (e Config) Ping() error {
	res, err := client.Ping()
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return errors.New(res.String())
	}

	return nil
}

// Reset elasticsearch client connection
func (e Config) Reset() {
	once.Reset()
}
