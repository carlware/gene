package utils

import (
	"bytes"
	"html/template"
)

func EvalString(tmpl string, doc interface{}) ([]byte, error) {
	tpl, err := template.New("test").Parse(tmpl)
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
