package attrs

import (
	"strings"

	"github.com/gobuffalo/flect"
	"github.com/gobuffalo/flect/name"
)

type Attr struct {
	Original  string
	Name      name.Ident
	inputType string
	goType    string
}
type Attrs []Attr

func (a Attr) String() string {
	return a.Original
}

func (ats Attrs) Slice() []string {
	var x []string
	for _, a := range ats {
		x = append(x, a.Original)
	}
	return x
}

// TODO need test case
// Supported Attr.inputType
// bigint, int, decimal, float, nulls.int,  nulls.bigint
// string, text, nulls.string, nulls.text
// timestamp, datetime, date, time, nulls.*
// uuid, nulls.uuid
// bool, nulls.bool
// jsonb, json
// blob
// slices.int, slices.float, slices.string

//GoType returns the Go type for an Attr based on its commonType
func (a Attr) GoType() string {
	if a.goType != "" {
		return a.goType
	}
	switch strings.ToLower(a.inputType) {
	//numeric, other numeric will be original name
	case "bigint":
		return "int64"
	case "decimal", "float":
		return "float64"
	case "nulls.bigint":
		return "nulls.Int64"
	case "nulls.int":
		return "nulls.Int"
	case "nulls.float":
		return "nulls.Float64"
	// text
	case "text":
		return "string"
	case "nulls.string", "nulls.text":
		return "nulls.String"
	// time
	case "timestamp", "datetime", "date", "time":
		return "time.Time"
	case "nulls.timestamp", "nulls.datetime", "nulls.date", "nulls.time":
		return "nulls.Time"
	// uuid
	case "uuid":
		return "uuid.UUID"
	case "nulls.uuid":
		return "nulls.UUID"
	//bool
	case "nulls.bool":
		return "nulls.Bool"
	case "json", "jsonb":
		return "slices.Map"
	// slices
	case "slices.string":
		return "slices.String"
	case "slices.int":
		return "slices.Int"
	case "slices.float", "slices.float32", "slices.float64":
		return "slices.Float"
	case "blob":
		return "[]byte"
	}
	return a.inputType
}

//GraphqlType return the type for graphql schema
func (a Attr) GraphqlType() string {
	return graphqlType(a.inputType, false)
}
func withNullable(s string, nullable bool) string {
	if nullable {
		return s
	} else {
		return s + "!"
	}
}
func graphqlType(s string, nullable bool) string {
	switch strings.ToLower(s) {
	case "int":
		return withNullable("Int", nullable)
	case "bigint":
		return withNullable("Int64", nullable)
	case "decimal", "float":
		return withNullable("Float", nullable)
	case "timestamp", "datetime", "date", "time":
		return withNullable("Time", nullable)
	case "text", "string":
		return withNullable("String", nullable)
	case "uuid.uuid", "uuid":
		return withNullable("UUID", nullable)
	case "slices.string":
		return "[String!]"
	case "slices.float":
		return "[Float!]"
	case "slices.int":
		return "[Int!]"
	case "json", "jsonb":
		return "Map!"
	case "blob":
		return "Any!"
	default:
		if strings.HasPrefix(s, "nulls.") {
			return graphqlType(strings.Replace(s, "nulls.", "", -1), true)
		}
		return flect.Pascalize(s)
	}
}

//FizzType returns the  type of database migration,
func (a Attr) FizzType() string {
	return fizzColType(a.inputType)
}

func fizzColType(s string) string {
	switch strings.ToLower(s) {
	case "int":
		return "integer"
	case "timestamp", "datetime", "date", "time":
		return "timestamp"
	case "uuid.uuid", "uuid":
		return "uuid"
	case "nulls.float32", "nulls.float64":
		return "float"
	case "slices.string":
		return "varchar[]"
	case "slices.float32", "slices.float64":
		return "numeric[]"
	case "json", "jsonb":
		return "jsonb"
	case "float32", "float64", "float":
		return "decimal"
	case "blob":
		return "blob"
	default:
		if strings.HasPrefix(s, "nulls.") {
			return fizzColType(strings.Replace(s, "nulls.", "", -1))
		}
		return strings.ToLower(s)
	}
}
