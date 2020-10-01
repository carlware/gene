package models

type Template struct {
	Path   string `yaml:"path"`
	Suffix string `yaml:"suffix"`
}

type Field map[string]interface{}

type Model struct {
	Name   string  `yaml:"name"`
	Plural string  `yaml:"plural"`
	Fields []Field `yaml:"fields"`
}

type Operation struct {
	Name     string `yaml:"name"`
	Filename string `yaml:"filename"`
	Template string `yaml:"template"`
}

type Generator struct {
	Operations []Operation `yaml:"operations"`
}

type Document struct {
	Root      string    `yaml:"root"`
	Template  Template  `yaml:"template"`
	Model     Model     `yaml:"model"`
	Generator Generator `yaml:"generator"`
}
