package neo4j

import (
	"fmt"

	"github.com/matryer/resync"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

var (
	driver neo4j.Driver
	err    error
	once   resync.Once
)

// Config struct to create new neo4j connection client
//
// Address: the neo4j database name
// Username: the neo4j database username
// Password: the neo4j database password
// LogEnabled: sets true if enable database log.
// MaxConnPool: sets the maximum number of open connections to the database.
// LogLevel: sets log mode, 1(Error) - 2(Warning) - 3(Info) - 4(Debug), default is Error
type Config struct {
	Address     string
	Username    string
	Password    string
	MaxConnPool int
	LogEnabled  bool
	LogLevel    int
}

// Driver create new neo4j connection driver
func (n Config) Driver() (neo4j.Driver, error) {
	url := fmt.Sprintf("%s", n.Address)
	auth := neo4j.BasicAuth(n.Username, n.Password, "")

	once.Do(func() {
		driver, err = neo4j.NewDriver(url, auth, func(config *neo4j.Config) {
			config.MaxConnectionPoolSize = n.MaxConnPool
			if n.LogEnabled {
				config.Log = neo4j.ConsoleLogger(neo4j.LogLevel(n.LogLevel))
			}
		})
	})

	return driver, err
}

// Reset reset neo4j connection driver
func (n Config) Reset() {
	once.Reset()
}
