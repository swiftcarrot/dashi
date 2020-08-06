package attrs

import (
	"strings"

	"github.com/gobuffalo/nulls"
	"github.com/swiftcarrot/dashi/generators/attrs/database"
	"github.com/swiftcarrot/flect"
)

type Attr struct {
	Original  string
	Name      flect.Ident
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
	case "integer":
		return "int"
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
	case "boolean":
		return "bool"
	case "nulls.boolean":
		return "nulls.Bool"
	case "json", "jsonb":
		return "slices.Map"
	case "blob":
		return "[]byte"
	case "strings":
		return "slices.String"
	case "integers":
		return "slices.Int"
	case "floats":
		return "slices.Float"
	case "uuids":
		return "slices.UUID"
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
	case "int", "integer":
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
	case "boolean":
		return withNullable("Boolean", nullable)
	case "strings":
		return "[String!]"
	case "floats":
		return "[Float!]"
	case "integers":
		return "[Int!]"
	case "uuids":
		return "[UUID!]"
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

func postgresType(s string) string {
	switch strings.ToLower(s) {
	case "int", "integer":
		return "int"
	case "timestamp", "datetime", "time":
		return "timestamp"
	case "text", "string":
		return "text"
	case "uuid.uuid", "uuid":
		return "uuid"
	case "boolean":
		return "boolean"
	case "strings":
		return "_text"
	case "floats":
		return "_float8"
	case "integers":
		return "_int4"
	case "uuids":
		return "_uuid"
	default:
		if strings.HasPrefix(s, "nulls.") {
			return postgresType(strings.Replace(s, "nulls.", "", -1))
		}
		return flect.Underscore(s)
	}
}

func (a Attr) PostgresColumn() database.Column {
	isSeq := false
	primary := false
	seqSuffix := ""
	colType := postgresType(a.inputType)
	defaultValue := nulls.String{
		Valid: false,
	}
	if a.Name.String() == "id" {
		primary = true
	}
	if a.Name.String() == "id" && colType != "uuid" {
		isSeq = true
		seqSuffix = "_" + a.Name.String() + "_seq"
	}
	return database.Column{
		Name:           a.Name,
		IsSequence:     isSeq,
		Default:        defaultValue,
		ColType:        postgresType(a.inputType),
		Nullable:       strings.HasPrefix(a.inputType, "nulls."),
		Primary:        primary,
		SequenceSuffix: seqSuffix,
	}

}

//FizzType returns the  type of database migration,
func (a Attr) FizzType() string {
	return fizzColType(a.inputType)
}

func fizzColType(s string) string {
	switch strings.ToLower(s) {
	case "int", "integer":
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
