package graphql

import (
	"encoding/json"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gobuffalo/nulls"
)

func MarshalNullsInt(s nulls.Int) graphql.Marshaler {
	if s.Valid {
		return graphql.MarshalInt(s.Int)
	} else {
		return graphql.MarshalAny(nil)
	}
}

func UnmarshalNullsInt(v interface{}) (nulls.Int, error) {
	switch v := v.(type) {
	case string, json.Number, int64, int:
		f, _ := graphql.UnmarshalInt(v)
		return nulls.NewInt(f), nil
	case nil:
		return nulls.Int{Valid: false, Int: 0}, nil
	default:
		return nulls.Int{Valid: false, Int: 0}, fmt.Errorf("%T is not a valid float", v)
	}
}

func MarshalNullsInt32(s nulls.Int32) graphql.Marshaler {
	if s.Valid {
		return graphql.MarshalInt32(s.Int32)
	} else {
		return graphql.MarshalAny(nil)
	}
}

func UnmarshalNullsInt32(v interface{}) (nulls.Int32, error) {
	switch v := v.(type) {
	case string, json.Number, int64, int:
		f, _ := graphql.UnmarshalInt32(v)
		return nulls.NewInt32(f), nil
	case nil:
		return nulls.Int32{Valid: false, Int32: 0}, nil
	default:
		return nulls.Int32{Valid: false, Int32: 0}, fmt.Errorf("%T is not a valid float", v)
	}
}

func MarshalNullsInt64(s nulls.Int64) graphql.Marshaler {
	if s.Valid {
		return graphql.MarshalInt64(s.Int64)
	} else {
		return graphql.MarshalAny(nil)
	}
}

func UnmarshalNullsInt64(v interface{}) (nulls.Int64, error) {
	switch v := v.(type) {
	case string, json.Number, int64, int:
		f, _ := graphql.UnmarshalInt64(v)
		return nulls.NewInt64(f), nil
	case nil:
		return nulls.Int64{Valid: false, Int64: 0}, nil
	default:
		return nulls.Int64{Valid: false, Int64: 0}, fmt.Errorf("%T is not a valid float", v)
	}
}
