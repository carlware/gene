package generator

import (
	"bytes"
	"carlware/gene/internal/models"
	"carlware/gene/internal/utils"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"

	"fmt"

	"gopkg.in/yaml.v2"
)

var funcs = template.FuncMap{"title": strings.Title}

func Generate(operation interface{}, temp, output string) error {
	tmplString, err := utils.FileToString(temp)
	if err != nil {
		fmt.Println("error", err)
	}

	tmpl, err := template.New("test").Funcs(funcs).Parse(tmplString)
	if err != nil {
		fmt.Println("error", err)
	}
	out := &bytes.Buffer{}

	err = tmpl.Execute(out, operation)
	if err != nil {
		fmt.Println("err", err)
	}

	err = utils.CreateFolders(output)
	if err != nil {
		fmt.Println("cannot create folder", err)
	}

	err = ioutil.WriteFile(output, out.Bytes(), os.ModePerm)
	if err != nil {
		fmt.Println("err", err)
	}

	return nil
}

func RunOperation(oper *models.Operator) {
	fmt.Println(oper)
	fmt.Println("model name", oper.ModelName)
	Generate(oper, oper.Template, oper.Path)

}

func ParseTemplate(templatePath string) (*models.Document, error) {
	tplString, err := ioutil.ReadFile(templatePath)
	if err != nil {
		return nil, err
	}

	doc := &models.Document{}
	err = yaml.Unmarshal(tplString, doc)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func ExcludeKeys(arr []models.Field, keys []string) []models.Field {
	fields := []models.Field{}
	for _, field := range arr {
		name, ok := field["name"]
		if !ok {
			continue
		}
		excluded := false
		for _, key := range keys {
			if name == key {
				excluded = true
				break
			}
		}
		if !excluded {
			fields = append(fields, field)
		}
	}
	return fields
}

func Gen(templatePath string) {
	doc, err := ParseTemplate(templatePath)
	if err != nil {
		fmt.Println("template error", err)
	}

	// masterTmpl, err := template.New("master").Funcs(funcs)

	fmt.Println("operations")
	for _, oper := range doc.Generator.Operations {
		fmt.Println("#### operation ###")
		op := &models.Operator{
			Name:        oper.Name,
			Template:    path.Join(doc.TemplatePath, oper.Template),
			ModelPlural: doc.Model.Plural,
			ModelName:   doc.Model.Name,
			Fields:      ExcludeKeys(doc.Model.Fields, oper.ExcludeFields),
			Properties:  oper.Properties,
		}

		// TODO: verify error
		path, err := utils.EvalString(oper.Path, op)
		if err != nil {
			panic(err)
		}

		op.Path = doc.Root + string(path)

		RunOperation(op)
		fmt.Println("#############")
	}

}
