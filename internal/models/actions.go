package models

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
	Request    Request    `yaml:"request"`
	Response   Response   `yaml:"response"`
}
