package graphql

import (
	"encoding/json"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gobuffalo/nulls"
)

func MarshalNullsFloat(s nulls.Float64) graphql.Marshaler {
	if s.Valid {
		return graphql.MarshalFloat(s.Float64)
	} else {
		return graphql.MarshalAny(nil)
	}
}

func UnmarshalNullsFloat(v interface{}) (nulls.Float64, error) {
	switch v := v.(type) {
	case string, json.Number, int64, int, float64:
		f, _ := graphql.UnmarshalFloat(v)
		return nulls.NewFloat64(f), nil
	case nil:
		return nulls.Float64{Valid: false, Float64: 0}, nil
	default:
		return nulls.Float64{Valid: false, Float64: 0}, fmt.Errorf("%T is not a valid float", v)
	}
}
