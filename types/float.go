package types

import (
	"encoding/json"
	"fmt"

	"github.com/gobuffalo/nulls"
	"github.com/swiftcarrot/gqlgen/graphql"
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
		return nulls.Float64{Valid: false}, nil
	default:
		return nulls.Float64{Valid: false}, fmt.Errorf("%T is not a valid float", v)
	}
}

func NewFloat(f *float64) nulls.Float64 {
	if f != nil {
		return nulls.NewFloat64(*f)
	} else {
		return nulls.Float64{Valid: false}
	}
}
