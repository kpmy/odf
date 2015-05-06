package odf

import (
	"odf/generators"
	"odf/mappers"
	"odf/model"
	_ "odf/model/stub"
	"odf/xmlns"
	"os"
	"testing"
)

func TestModel(t *testing.T) {
	m := model.ModelFactory()
	if m == nil {
		t.Error("model is nil")
	}
	w := m.NewWriter()
	if w == nil {
		t.Error("writer is nil")
	}
}

func TestMappers(t *testing.T) {
	m := model.ModelFactory()
	fm := &mappers.Formatter{}
	fm.ConnectTo(m)
	fm.MimeType = xmlns.MimeText
	fm.Init()
}

func TestGenerators(t *testing.T) {
	output, _ := os.OpenFile("test0.odf", os.O_CREATE|os.O_WRONLY, 0666)
	m := model.ModelFactory()
	fm := &mappers.Formatter{}
	fm.ConnectTo(m)
	fm.MimeType = xmlns.MimeText
	fm.Init()
	generators.Generate(m, output, fm.MimeType)
	output.Close()
}

func TestStructure(t *testing.T) {
	output, _ := os.OpenFile("test0.odf", os.O_CREATE|os.O_WRONLY, 0666)
	m := model.ModelFactory()
	fm := &mappers.Formatter{}
	fm.ConnectTo(m)
	fm.MimeType = xmlns.MimeText
	fm.Init()
	fm.WriteString(`Hello, World!`)
	generators.Generate(m, output, fm.MimeType)
	output.Close()
}
