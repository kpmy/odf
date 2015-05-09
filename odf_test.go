package odf

import (
	"github.com/kpmy/ypk/assert"
	"image/color"
	"odf/generators"
	"odf/mappers"
	"odf/mappers/attr"
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
	generators.GeneratePackage(m, nil, output, fm.MimeType)
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
	generators.GeneratePackage(m, nil, output, fm.MimeType)
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
	fm.WriteString("Hello, World!\n")
	fm.SetAttr(new(attr.TextAttributes).Size(32).FontFace("Arial"))
	fm.WriteString(`Hello, Go!`)
	fm.SetAttr(new(attr.TextAttributes).Size(36).FontFace("Courier New").Bold().Italic())
	fm.WriteString(`	Hello, Again!`)
	fm.SetAttr(new(attr.TextAttributes).Size(32).FontFace("Arial")) //test attribute cache
	fm.SetAttr(new(attr.TextAttributes).Size(32).FontFace("Arial").Color(color.RGBA{0x00, 0xff, 0xff, 0xff}))
	fm.WriteString("\nNo, not you again!")
	fm.SetAttr(new(attr.ParagraphAttributes).AlignRight().PageBreak())
	fm.WritePara("Page break!\r")
	fm.SetAttr(nil)
	fm.WriteString(`Hello, Пщ!`)
	generators.GeneratePackage(m, nil, output, fm.MimeType)
	assert.For(output.Close() == nil, 20)
}

func TestTables(t *testing.T) {
	{
		output, _ := os.OpenFile("test3.odf", os.O_CREATE|os.O_WRONLY, 0666)
		m := model.ModelFactory()
		fm := &mappers.Formatter{}
		fm.ConnectTo(m)
		fm.MimeType = xmlns.MimeText
		fm.Init()
		generators.GeneratePackage(m, nil, output, fm.MimeType)
		assert.For(output.Close() == nil, 20)
	}
	{
		output, _ := os.OpenFile("test4.odf", os.O_CREATE|os.O_WRONLY, 0666)
		m := model.ModelFactory()
		fm := &mappers.Formatter{}
		fm.ConnectTo(m)
		fm.MimeType = xmlns.MimeSpreadsheet
		fm.Init()
		generators.GeneratePackage(m, nil, output, fm.MimeType)
		assert.For(output.Close() == nil, 20)
	}
}

func TestDraw(t *testing.T) {
	const ImagePng xmlns.Mime = "image/png"
	output, _ := os.OpenFile("test5.odf", os.O_CREATE|os.O_WRONLY, 0666)
	m := model.ModelFactory()
	fm := &mappers.Formatter{}
	fm.ConnectTo(m)
	fm.MimeType = xmlns.MimeText
	fm.Init()
	embed := make(map[string]generators.Embeddable)
	{
		img, _ := os.Open("project.png")
		d := mappers.NewDraw(img, ImagePng)
		url := d.WriteTo(fm, "Two Gophers", 6.07, 3.53) //magic? real size of `project.png`
		embed[url] = d
	}
	generators.GeneratePackage(m, embed, output, fm.MimeType)
	assert.For(output.Close() == nil, 20)
}
