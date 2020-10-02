package generator

import (
	"bytes"
	"carlware/gene/internal/funcs"
	"carlware/gene/internal/models"
	"carlware/gene/internal/utils"
	"io/ioutil"
	"path"
	"text/template"

	"fmt"

	"gopkg.in/yaml.v2"
)

func Generate(operation interface{}, templateName, temp, output string) error {
	tmplString, err := utils.FileToString(temp)
	if err != nil {
		return err
	}

	tmpl, err := template.New(templateName).Funcs(funcs.Funcs).Parse(tmplString)
	if err != nil {
		return err
	}
	out := &bytes.Buffer{}

	err = tmpl.Execute(out, operation)
	if err != nil {
		return err
	}

	err = utils.BytesToFile(output, out.Bytes())
	if err != nil {
		return err
	}

	return nil
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
		fmt.Println("error while parsing template")
		panic(err)
	}

	for _, oper := range doc.Generator.Operations {
		op := &models.Gene{
			Name:     oper.Name,
			Template: path.Join(doc.TemplatesPath, oper.Template),
			Model: &models.Model{
				Name:   doc.Model.Name,
				Plural: doc.Model.Plural,
				Fields: ExcludeKeys(doc.Model.Fields, oper.ExcludeFields),
			},
			Properties: oper.Properties,
		}

		dst, err := utils.EvalString(oper.Path, op)
		if err != nil {
			fmt.Println("error while evaluating the path", oper.Path)
			panic(err)
		}
		op.Path = path.Join(doc.Destination, string(dst))

		err = Generate(op, op.Name, op.Template, op.Path)
		if err != nil {
			fmt.Println("error while generating files")
			panic(err)
		}
	}

	fmt.Println("files generated successfully")

}
