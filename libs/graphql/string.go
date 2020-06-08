package graphql

import (
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gobuffalo/nulls"
)

func MarshalString(s string) graphql.Marshaler {
	return graphql.MarshalString(s)
}

func UnmarshalString(v interface{}) (string, error) {
	return graphql.UnmarshalString(v)
}

func MarshalNullsString(s nulls.String) graphql.Marshaler {
	if s.Valid {
		return graphql.MarshalString(s.String)
	} else {
		return graphql.MarshalAny(nil)
	}
}

func UnmarshalNullsString(v interface{}) (nulls.String, error) {
	switch v := v.(type) {
	case string:
		return nulls.NewString(v), nil
	case nil:
		return nulls.String{Valid: false}, nil
	default:
		return nulls.String{Valid: false}, fmt.Errorf("%T is not a valid uuid", v)
	}
}

func NewString(f *string) nulls.String {
	if f != nil {
		return nulls.NewString(*f)
	} else {
		return nulls.String{Valid: false}
	}
}
