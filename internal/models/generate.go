package models

type Field map[string]interface{}
type Properties map[string]interface{}

type Base struct {
	Name       string     `yaml:"name"`
	Plural     string     `yaml:"plural"`
	Properties Properties `yaml:"properties"`
	Fields     []Field    `yaml:"fields"`
}

type Request struct {
	Properties Properties `yaml:"properties"`
	Fields     []Field    `yaml:"fields"`
}

type Response struct {
	Properties Properties `yaml:"properties"`
	Fields     []Field    `yaml:"fields"`
}

type Action struct {
	Name       string     `yaml:"name"`
	Plural     string     `yaml:"plural"`
	Properties Properties `yaml:"properties"`
	Inherit    bool       `yaml:"inherit"`
	Exclude    []string   `yaml:"exclude"`
	Request    *Request   `yaml:"request"`
	Response   *Response  `yaml:"response"`
}

type Model struct {
	Base `yaml:",inline"`
}

type Operation struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`

	// operation types
	Generic Generic `yaml:"generic"`
	Process Process `yaml:"process"`
}

type Document struct {
	Destination   string      `yaml:"destination"`
	TemplatesPath string      `yaml:"templates_path"`
	Model         Model       `yaml:"model"`
	Actions       []Action    `yaml:"actions"`
	Operations    []Operation `yaml:"operations"`
}
