package models

type Process struct {
	Path           string   `yaml:"path"`
	Filename       string   `yaml:"filename"`
	Template       string   `yaml:"template"`
	ExcludeActions []string `yaml:"exclude_actions,flow"`

	// This properties will be passed at runtime
	Model   Model
	Actions []Action
	Action  *Action
	Name    string
}
