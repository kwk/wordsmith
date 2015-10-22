package main

import (
	"errors"
	"io"
	"text/template"
)

type Format uint

const (
	JSON Format = iota
)

type Metadata struct {
	PackageName   string
	Object        string
	MarshalObject string
	Type          string
}

type Generator struct {
	Format Format
}

func (g *Generator) Generate(writer io.Writer, metadata Metadata) error {
	tmpl, err := g.template()
	if err != nil {
		return nil
	}

	return tmpl.Execute(writer, metadata)
}

func (g *Generator) template() (*template.Template, error) {
	if g.Format != JSON {
		return nil, errors.New("Unsupported format")
	}

	resource, err := Asset("templates/writeto_json.tmpl")
	if err != nil {
		return nil, err
	}

	tmpl := template.New("template")
	return tmpl.Parse(string(resource))
}
