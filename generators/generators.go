package generators

type GenerateSchema struct {
	Name   string
	Fields *[]Field
}

type Field struct {
	Name string
	Type string
}

func GenerateScaffold(schema *GenerateSchema) error {
	return nil
}
