package main

import (
	"bytes"
	"io"
	"odf/generators"
	"odf/mappers"
	"odf/model"
	"odf/xmlns"
)

func report() (io.Reader, error) {
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
