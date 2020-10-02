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
	Name          string                 `yaml:"name"`
	Filename      string                 `yaml:"filename"`
	Template      string                 `yaml:"template"`
	Path          string                 `yaml:"path"`
	ExcludeFields []string               `yaml:"exclude_fields"`
	Properties    map[string]interface{} `yaml:"properties"`
}

type Generator struct {
	Operations []Operation `yaml:"operations"`
}

type Document struct {
	Root         string    `yaml:"root"`
	TemplatePath string    `yaml:"template_path"`
	Model        Model     `yaml:"model"`
	Generator    Generator `yaml:"generator"`
}

type Operator struct {
	Name        string                 `yaml:"name"`
	Model       Model                  `yaml:"model"`
	ModelName   string                 `yaml:"model_name"`
	ModelPlural string                 `yaml:"model_plural"`
	Template    string                 `yaml:"template"`
	Suffix      string                 `yaml:"suffix"`
	Path        string                 `yaml:"path"`
	Fields      []Field                `yaml:"fields"`
	Properties  map[string]interface{} `yaml:"properties"`
}
