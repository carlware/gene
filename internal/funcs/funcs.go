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

func has(in []interface{}, key string) bool {
	for _, param := range in {
		if val, ok := param.(string); ok {
			if val == key {
				return true
			}
		}
	}
	return false
}

var Funcs = template.FuncMap{
	"title":  strings.Title,
	"join":   strings.Join,
	"upper":  strings.ToUpper,
	"hyphen": hyphen,
	"has":    has,
}
