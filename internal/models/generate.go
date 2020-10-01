package models

type Template struct {
	Path   string
	Suffix string
}

type Field struct {
	Name string
	Type string
}

type Model struct {
	Name   string
	Fields []Field
}

type Operation struct {
	Name     string
	Filename string
	Template string
}

type Generator struct {
	Operations []Operation
}

type Document struct {
	Root     string
	Template Template
	Model    Model
	Generate Generate
}
