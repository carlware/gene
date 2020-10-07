package generator

import (
	"carlware/gene/internal/models"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

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

func ExcludeActions(arr []models.Action, keys []string) []models.Action {
	actions := []models.Action{}
	for _, action := range arr {
		excluded := false
		for _, key := range keys {
			if action.Name == key {
				excluded = true
				break
			}
		}
		if !excluded {
			actions = append(actions, action)
		}
	}
	return actions
}
