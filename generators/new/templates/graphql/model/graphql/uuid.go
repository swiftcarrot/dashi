package graphql

import (
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gobuffalo/nulls"
	"github.com/gofrs/uuid"
)

func MarshalUUID(u uuid.UUID) graphql.Marshaler {
	return graphql.MarshalString(u.String())
}

func UnmarshalUUID(v interface{}) (uuid.UUID, error) {
	switch v := v.(type) {
	case string:
		return uuid.FromStringOrNil(v), nil
	default:
		return uuid.Nil, fmt.Errorf("%T is not a valid uuid", v)
	}
}

func MarshalNullsUUID(u nulls.UUID) graphql.Marshaler {
	if u.Valid {
		return graphql.MarshalString(u.UUID.String())
	} else {
		return graphql.MarshalAny(nil)
	}
}

func UnmarshalNullsUUID(v interface{}) (nulls.UUID, error) {
	switch v := v.(type) {
	case string:
		return nulls.NewUUID(uuid.FromStringOrNil(v)), nil
	case nil:
		return nulls.UUID{
			UUID:  uuid.Nil,
			Valid: false,
		}, nil
	default:
		return nulls.UUID{
			UUID:  uuid.Nil,
			Valid: false,
		}, fmt.Errorf("%T is not a valid uuid", v)
	}
}
