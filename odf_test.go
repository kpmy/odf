package odf

import (
	"odf/generators"
	"odf/mappers"
	"odf/mappers/attr"
	"odf/model"
	_ "odf/model/stub"
	"odf/xmlns"
	"os"
	"testing"
	"ypk/assert"
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
	assert.For(output.Close() == nil, 20)
}

func TestStructure(t *testing.T) {
	output, _ := os.OpenFile("test1.odf", os.O_CREATE|os.O_WRONLY, 0666)
	m := model.ModelFactory()
	fm := &mappers.Formatter{}
	fm.ConnectTo(m)
	fm.MimeType = xmlns.MimeText
	fm.Init()
	fm.WriteString("Hello, World!   \t   \n   \r	фыва 	фыва		\n фыва")
	generators.Generate(m, output, fm.MimeType)
	assert.For(output.Close() == nil, 20)
}

func TestStylesMechanism(t *testing.T) {
	output, _ := os.OpenFile("test2.odf", os.O_CREATE|os.O_WRONLY, 0666)
	m := model.ModelFactory()
	fm := &mappers.Formatter{}
	fm.ConnectTo(m)
	fm.MimeType = xmlns.MimeText
	fm.Init()
	fm.RegisterFont("Arial", "Arial")
	fm.RegisterFont("Courier New", "Courier New")
	fm.SetDefaults(new(attr.TextAttributes).Size(18).FontFace("Courier New"))
	fm.SetDefaults(new(attr.TextAttributes).Size(16).FontFace("Courier New"))
	fm.WriteString(`Hello, World!`)
	fm.SetAttr(new(attr.TextAttributes).Size(32).FontFace("Arial"))
	fm.WriteString(`Hello, World!`)
	fm.SetAttr(new(attr.TextAttributes).Size(36).FontFace("Courier New"))
	fm.WriteString(`Hello, World!`)
	fm.SetAttr(new(attr.TextAttributes).Size(32).FontFace("Arial")) //test attribute cache
	fm.WriteString(`Hello, World!`)
	fm.SetAttr(new(attr.ParagraphAttributes).AlignRight().PageBreak())
	fm.WritePara(`Page break!`)
	fm.SetAttr(nil)
	fm.WriteString(`Hello, Пщ!`)
	generators.Generate(m, output, fm.MimeType)
	assert.For(output.Close() == nil, 20)
}

func TestTables(t *testing.T) {
	table := func(fm *mappers.Formatter) {
		tm := &mappers.TableMapper{}
		tm.ConnectTo(fm)
		tm.Write("test", 5, 10)
		tt := tm.List["test"]
		tm.WriteColumns(tt, 4)
		tm.WriteRows(tt, 3)
	}
	{
		output, _ := os.OpenFile("test3.odf", os.O_CREATE|os.O_WRONLY, 0666)
		m := model.ModelFactory()
		fm := &mappers.Formatter{}
		fm.ConnectTo(m)
		fm.MimeType = xmlns.MimeText
		fm.Init()
		table(fm)
		generators.Generate(m, output, fm.MimeType)
		assert.For(output.Close() == nil, 20)
	}
	{
		output, _ := os.OpenFile("test4.odf", os.O_CREATE|os.O_WRONLY, 0666)
		m := model.ModelFactory()
		fm := &mappers.Formatter{}
		fm.ConnectTo(m)
		fm.MimeType = xmlns.MimeSpreadsheet
		fm.Init()
		table(fm)
		generators.Generate(m, output, fm.MimeType)
		assert.For(output.Close() == nil, 20)
	}
}
