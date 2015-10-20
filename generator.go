package main

import (
	"fmt"
	"io"
)

var JSONTemplate string

type Format uint

const (
	JSON Format = iota
)

type Generator struct {
	Type   string
	Format Format
}

func (g *Generator) Generate(writer io.Writer) error {
	var template string

	if g.Format == JSON {
		template = JSONTemplate
	}

	_, err := fmt.Fprintf(writer, template, g.Type)
	return err
}
