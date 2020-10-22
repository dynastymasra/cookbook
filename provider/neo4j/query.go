package neo4j

const (
	Equal = "Equal"
	In    = "In"

	Descending = "Descending"
	Ascending  = "Ascending"
)

var (
	validOrdering = map[string]bool{
		Descending: true,
		Ascending:  true,
	}
)

type (
	Filter struct {
		Condition string
		Field     string
		Value     interface{}
	}

	Ordering struct {
		Field     string
		Direction string
	}

	Query struct {
		Node      string
		Limit     int
		Offset    int
		Filters   []*Filter
		Orderings []*Ordering
	}
)

func NewQuery(node string) *Query {
	return &Query{
		Node: node,
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

// Order adds a sort order to the query
func (q *Query) Ordering(property, direction string) *Query {
	order := NewOrdering(property, direction)
	q.Orderings = append(q.Orderings, order)
	return q
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
