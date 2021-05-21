package dashboardgen

import (
	"github.com/swiftcarrot/dashi/gqlgen/codegen"
	"github.com/swiftcarrot/dashi/gqlgen/plugin"
)

type Plugin struct {
	ObjectName string
}

type DashboardData struct {
	Object    *codegen.Object
	Resolvers []*Resolver
}

type Resolver struct {
	Object *codegen.Object
	Field  *codegen.Field
}

var _ plugin.CodeGenerator = &Plugin{}

func New(obj string) plugin.Plugin {
	return &Plugin{
		ObjectName: obj,
	}
}

func (m *Plugin) Name() string {
	return "dashboardgen"
}

func (m *Plugin) GenerateCode(data *codegen.Data) error {
	return m.generatePerSchema(data)
}

// TODO extract entity and query/mutation from codegen.Data, and generate corresponding js api code
func (m *Plugin) generatePerSchema(data *codegen.Data) error {
	for _, o := range data.Objects {
		resolvers := []*Resolver{}
		if o.Name == m.ObjectName {
			for _, f := range o.Fields {
				if !f.IsResolver {
					continue
				} else {
					resolvers = append(resolvers, &Resolver{
						Object: o,
						Field:  f,
					})
				}
			}
			break
		}
	}
	return nil
}
