package storybook

import (
	"encoding/json"
	"os/exec"
	"text/template"

	"github.com/gobuffalo/genny/v2"
	"github.com/gobuffalo/genny/v2/gogen"
	"github.com/gobuffalo/packr/v2"
)

func New(opts *Options) (*genny.Generator, error) {
	g := genny.New()

	if err := opts.Validate(); err != nil {
		return g, err
	}

	if err := g.Box(packr.New("storybook/templates", "../storybook/templates")); err != nil {
		return g, err
	}

	data := map[string]interface{}{
		"opts": opts,
	}
	helpers := template.FuncMap{}
	t := gogen.TemplateTransformer(data, helpers)
	g.Transformer(t)
	g.Transformer(genny.Dot())
	g.Command(exec.Command("yarn", "add", "@babel/core", "babel-loader", "babel-preset-swiftcarrot", "@storybook/addon-actions", "@storybook/addon-links", "@storybook/addons", "@storybook/react", "glob", "--dev", "-W"))

	// TODO: add storybook scripts to package.json
	g.RunFn(func(r *genny.Runner) error {
		f, err := r.Disk.Find("package.json")
		if err != nil {
			return err
		}

		packageJSON := make(map[string]interface{})
		err = json.Unmarshal([]byte(f.String()), &packageJSON)
		if err != nil {
			panic(err)
		}

		scripts := make(map[string]interface{})
		if packageJSON["scripts"] != nil {
			scripts = packageJSON["scripts"].(map[string]interface{})
		}

		scripts["storybook"] = "start-storybook -p 6006"
		scripts["build-storybook"] = "build-storybook"
		packageJSON["scripts"] = scripts

		// output, err := json.Marshal(packageJSON)
		// fmt.Println(string(output))

		return nil
	})

	return g, nil
}
