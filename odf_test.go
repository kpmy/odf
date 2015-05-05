package odf

import (
	"odf/mappers"
	"odf/model"
	_ "odf/model/stub"
	"odf/xmlns"
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
