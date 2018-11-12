package instrumentation

import "github.com/newrelic/go-agent"

// NewRelicApplication create new relic application
func NewRelicApplication(config newrelic.Config) (newrelic.Application, error) {
	app, err := newrelic.NewApplication(config)
	if err != nil {
		return nil, err
	}
	return app, nil
}
