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

	once.Do(func() {
		client, errs = elasticsearch.NewClient(config)
	})

	res, err := client.Ping()
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, errors.New(res.String())
	}

	return client, errs
}

// Reset elasticsearch client connection
func (e Config) Reset() {
	once.Reset()
}
