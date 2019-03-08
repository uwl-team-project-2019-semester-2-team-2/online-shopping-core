package database

type QueryBuilder struct {
	Query string
}

func (q *QueryBuilder) Select(query string) *QueryBuilder {


	return q
}