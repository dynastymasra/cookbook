package instrumentation

import (
	"context"
	"net/http"

	"github.com/newrelic/go-agent"
)

// NewRelicApplication create new relic application
func NewRelicApplication(config newrelic.Config) (newrelic.Application, error) {
	app, err := newrelic.NewApplication(config)
	if err != nil {
		return nil, err
	}
	return app, nil
}

// NewRelicDataStoreSegment used for send database data to new relic
// defer NewRelicDataStoreSegment(...)()
func NewRelicDataStoreSegment(ctx context.Context, product newrelic.DatastoreProduct, queryParams map[string]interface{},
	collection, operation, query, host, port, databaseName string) func() {
	if trx := newrelic.FromContext(ctx); trx != nil {
		segment := &newrelic.DatastoreSegment{
			StartTime:          newrelic.StartSegmentNow(trx),
			Product:            product,
			QueryParameters:    queryParams,
			Collection:         collection,
			Operation:          operation,
			ParameterizedQuery: query,
			Host:               host,
			PortPathOrID:       port,
			DatabaseName:       databaseName,
		}
		return func() { segment.End() }
	}
	return func() {}
}

// NewRelicExternalSegment used for send external data request to new relic
// segment := NewRelicExternalSegment(...)
// resp, err := client.Do(request)
// segment.Response = resp
// segment.End()
func NewRelicExternalSegment(ctx context.Context, request *http.Request, url string) *newrelic.ExternalSegment {
	if trx := newrelic.FromContext(ctx); trx != nil {
		return &newrelic.ExternalSegment{
			StartTime: newrelic.StartSegmentNow(trx),
			URL:       url,
			Request:   request,
		}
	}
	return &newrelic.ExternalSegment{}
}
