package mongo

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Query struct {
	Limit     int
	Offset    int
	Filters   []*Filter
	Orderings []*Ordering
}

type Filter struct {
	Condition string
	Field     string
	Value     interface{}
}

type Ordering struct {
	Field     string
	Direction string
}

const (
	Equal            = "Equal"
	LessThan         = "LessThan"
	LessThanEqual    = "LessThanEqual"
	GreaterThan      = "GreaterThan"
	GreaterThanEqual = "GreaterThanEqual"

	Descending = "Descending"
	Ascending  = "Ascending"
)

var (
	validOrdering = map[string]bool{
		Descending: true,
		Ascending:  true,
	}
)

func NewQuery(limit, offset int) *Query {
	return &Query{
		Limit:  limit,
		Offset: offset,
	}
}

func NewOrdering(field, direction string) *Ordering {
	d := direction

	if !validOrdering[direction] {
		d = Descending
	}

	return &Ordering{
		Field:     field,
		Direction: d,
	}
}

// NewFilter creates a new property filter
func NewFilter(field, condition string, value interface{}) *Filter {
	return &Filter{
		Field:     field,
		Condition: condition,
		Value:     value,
	}
}

// Filter adds a filter to the query
func (q *Query) Filter(property, condition string, value interface{}) *Query {
	filter := NewFilter(property, condition, value)
	q.Filters = append(q.Filters, filter)
	return q
}

func (q *Query) Slice(offset, limit int) *Query {
	q.Offset = offset
	q.Limit = limit

	return q
}

// Order adds a sort order to the query
func (q *Query) Ordering(property, direction string) *Query {
	order := NewOrdering(property, direction)
	q.Orderings = append(q.Orderings, order)
	return q
}

func TranslateQuery(query *Query) (bson.D, *options.FindOptions) {
	limit := int64(query.Limit)
	skip := int64(query.Offset)

	d := bson.D{}

	for _, filter := range query.Filters {
		switch filter.Condition {
		case Equal:
			d = append(d, bson.E{
				Key:   fmt.Sprintf("%v", filter.Field),
				Value: filter.Value,
			})
		case GreaterThan:
			d = append(d, bson.E{
				Key:   fmt.Sprintf("%v", filter.Field),
				Value: bson.M{"$gt": filter.Value},
			})
		case GreaterThanEqual:
			d = append(d, bson.E{
				Key:   fmt.Sprintf("%v", filter.Field),
				Value: bson.M{"$gte": filter.Value},
			})
		case LessThan:
			d = append(d, bson.E{
				Key:   fmt.Sprintf("%v", filter.Field),
				Value: bson.M{"$lt": filter.Value},
			})
		case LessThanEqual:
			d = append(d, bson.E{
				Key:   fmt.Sprintf("%v", filter.Field),
				Value: bson.M{"$lte": filter.Value},
			})
		default:
			d = append(d, bson.E{
				Key:   fmt.Sprintf("%v", filter.Field),
				Value: filter.Value,
			})
		}
	}

	for _, order := range query.Orderings {
		switch order.Direction {
		case Ascending:

		case Descending:

		default:

		}
	}

	option := &options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
	}

	return d, option
}
