package funcs

import (
	"strings"
	"text/template"
)

func hyphen(in string) string {
	newString := ""
	for _, char := range in {
		if string(char) >= "A" && string(char) <= "Z" {
			newString += ("-" + string(char+32))
		} else {
			newString += string(char)
		}
	}
	return newString
}

var Funcs = template.FuncMap{
	"title":  strings.Title,
	"join":   strings.Join,
	"upper":  strings.ToUpper,
	"hyphen": hyphen,
}
