package main

import (
	"bytes"
	"github.com/kpmy/odf/generators"
	"github.com/kpmy/odf/mappers"
	"github.com/kpmy/odf/model"
	"github.com/kpmy/odf/xmlns"
	"io"
)

func demo() (io.Reader, error) {
	output := bytes.NewBuffer(nil)
	m := model.ModelFactory()
	fm := &mappers.Formatter{}
	fm.ConnectTo(m)
	fm.MimeType = xmlns.MimeText
	fm.Init()
	fm.WriteString("Hello, World!")
	generators.GeneratePackage(m, nil, output, fm.MimeType)
	return output, nil
}
