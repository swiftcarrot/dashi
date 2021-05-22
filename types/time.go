package types

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/gobuffalo/nulls"
)

func MarshalNullsTime(s nulls.Time) graphql.Marshaler {
	if s.Valid {
		return graphql.MarshalTime(s.Time)
	} else {
		return graphql.Null
	}
}

func UnmarshalNullsTime(v interface{}) (nulls.Time, error) {
	var time, err = graphql.UnmarshalTime(v)
	if err != nil {
		return nulls.NewTime(time), nil
	}
	return nulls.Time{Valid: false}, err
}
