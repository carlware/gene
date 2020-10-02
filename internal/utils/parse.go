package utils

import (
	"bytes"
	"carlware/gene/internal/funcs"
	"text/template"
)

func EvalString(tmpl string, doc interface{}) ([]byte, error) {
	tpl, err := template.New("eval").Funcs(funcs.Funcs).Parse(tmpl)
	if err != nil {
		return []byte{}, err
	}
	out := &bytes.Buffer{}

	err = tpl.Execute(out, doc)
	if err != nil {
		return []byte{}, err
	}
	return out.Bytes(), err
}
