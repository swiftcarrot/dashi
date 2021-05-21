package stubgen

import (
	"path/filepath"
	"syscall"

	"github.com/swiftcarrot/dashi/gqlgen/internal/code"

	"github.com/swiftcarrot/dashi/gqlgen/codegen"
	"github.com/swiftcarrot/dashi/gqlgen/codegen/config"
	"github.com/swiftcarrot/dashi/gqlgen/codegen/templates"
	"github.com/swiftcarrot/dashi/gqlgen/plugin"
)

func New(filename string, typename string) plugin.Plugin {
	return &Plugin{filename: filename, typeName: typename}
}

type Plugin struct {
	filename string
	typeName string
}

var _ plugin.CodeGenerator = &Plugin{}
var _ plugin.ConfigMutator = &Plugin{}

func (m *Plugin) Name() string {
	return "stubgen"
}

func (m *Plugin) MutateConfig(cfg *config.Config) error {
	_ = syscall.Unlink(m.filename)
	return nil
}

func (m *Plugin) GenerateCode(data *codegen.Data) error {
	abs, err := filepath.Abs(m.filename)
	if err != nil {
		return err
	}
	pkgName := code.NameForDir(filepath.Dir(abs))

	return templates.Render(templates.Options{
		PackageName: pkgName,
		Filename:    m.filename,
		Data: &ResolverBuild{
			Data:     data,
			TypeName: m.typeName,
		},
		GeneratedHeader: true,
		Packages:        data.Config.Packages,
	})
}

type ResolverBuild struct {
	*codegen.Data

	TypeName string
}
