package slices

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
)

// Map is a map[string]interface.
type Map map[string]interface{}

// Interface implements the nulls.nullable interface.
func (m Map) Interface() interface{} {
	return map[string]interface{}(m)
}

// Scan implements the sql.Scanner interface.
// It allows to read the map from the database value.
func (m *Map) Scan(src interface{}) error {
	var b []byte
	switch t := src.(type) {
	case []byte:
		b = t
	case string:
		b = []byte(t)
	default:
		return fmt.Errorf("scan source was not []byte nor string but %T", src)
	}
	err := json.Unmarshal(b, m)
	if err != nil {
		return err
	}
	return nil
}

// Value implements the driver.Valuer interface.
// It allows to convert the map to a driver.value.
func (m Map) Value() (driver.Value, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return string(b), nil
}

// UnmarshalJSON will unmarshall JSON value into
// the map representation of this value.
func (m Map) UnmarshalJSON(b []byte) error {
	var stuff map[string]interface{}
	err := json.Unmarshal(b, &stuff)
	if err != nil {
		return err
	}
	for key, value := range stuff {
		m[key] = value
	}
	return nil
}

// UnmarshalText will unmarshall text value into
// the map representation of this value.
func (m Map) UnmarshalText(text []byte) error {
	err := json.Unmarshal(text, &m)
	if err != nil {
		return err
	}
	return nil
}

func MarshalMap(val Map) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		err := json.NewEncoder(w).Encode(val)
		if err != nil {
			panic(err)
		}
	})
}

func UnmarshalMap(v interface{}) (Map, error) {
	if m, ok := v.(Map); ok {
		return m, nil
	}

	return nil, fmt.Errorf("%T is not a map", v)
}