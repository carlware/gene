package models

type Field map[string]interface{}
type Properties map[string]interface{}

type Model struct {
	Name   string  `yaml:"name"`
	Plural string  `yaml:"plural"`
	Fields []Field `yaml:"fields"`
}

type Operation struct {
	Name          string     `yaml:"name"`
	Template      string     `yaml:"template"`
	Path          string     `yaml:"path"`
	ExcludeFields []string   `yaml:"exclude_fields"`
	Properties    Properties `yaml:"properties"`
}

type Generator struct {
	Operations []Operation `yaml:"operations"`
}

type Document struct {
	Destination   string    `yaml:"destination"`
	TemplatesPath string    `yaml:"templates_path"`
	Model         Model     `yaml:"model"`
	Generator     Generator `yaml:"generator"`
}

type Gene struct {
	Name       string     `yaml:"name"`
	Model      *Model     `yaml:"model"`
	Template   string     `yaml:"template"`
	Path       string     `yaml:"path"`
	Properties Properties `yaml:"properties"`
}
