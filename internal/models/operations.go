package models

type Generic struct {
	Template      string     `yaml:"template"`
	Path          string     `yaml:"path"`
	ExcludeFields []string   `yaml:"exclude_fields"`
	Properties    Properties `yaml:"properties"`
}

type Action struct {
	Name       string     `yaml:"name"`
	Template   string     `yaml:"template"`
	Properties Properties `yaml:"properties"`
	Fields     []Field    `yaml:"fields"`
}
