package mongo

import "go.mongodb.org/mongo-driver/mongo/options"

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
	Equal = "Equal"

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

func TranslateQuery(query *Query) (interface{}, *options.FindOptions) {
	limit := int64(query.Limit)
	skip := int64(query.Offset)

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

	return nil, option
}
