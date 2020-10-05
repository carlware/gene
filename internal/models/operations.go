package models

type Generic struct {
	Name          string     `yaml:"name"`
	Template      string     `yaml:"template"`
	Path          string     `yaml:"path"`
	ExcludeFields []string   `yaml:"exclude_fields"`
	Properties    Properties `yaml:"properties"`

	// This properties will be passed at runtime
	Model   *Model
	Actions []Action
}

type Process struct {
	Path     string   `yaml:"path"`
	Filename string   `yaml:"filename"`
	Template string   `yaml:"template"`
	Exclude  []string `yaml:"exclude,flow"`

	// This properties will be passed at runtime
	Model   Model
	Actions []Action
	Action  *Action
	Name    string
}
