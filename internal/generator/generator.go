package generator

import (
	"bytes"
	"carlware/gene/internal/funcs"
	"carlware/gene/internal/utils"
	"text/template"
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
