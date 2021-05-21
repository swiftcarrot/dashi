package types

import (
	"fmt"

	"github.com/gobuffalo/nulls"
	"github.com/gofrs/uuid"
	"github.com/swiftcarrot/dashi/gqlgen/graphql"
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
			Valid: false,
		}, nil
	default:
		return nulls.UUID{
			Valid: false,
		}, fmt.Errorf("%T is not a valid uuid", v)
	}
}

func NewUUID(f *uuid.UUID) nulls.UUID {
	if f != nil {
		return nulls.NewUUID(*f)
	} else {
		return nulls.UUID{Valid: false}
	}
}
