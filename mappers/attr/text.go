package attr

import (
	"image/color"
	"odf/model"
	"odf/xmlns/fo"
	"odf/xmlns/style"
	"odf/xmlns/text"
)

type TextAttributes struct {
	fontFace string
	size     int
	col      color.Color
	bold     bool
	italic   bool
	named
}

func (t *TextAttributes) Equal(_a Attributes) (ok bool) {
	a, ok := _a.(*TextAttributes)
	if ok {
		ok = t.size == a.size && t.fontFace == a.fontFace && t.col == a.col && t.italic == a.italic && t.bold == a.bold
	}
	return
}

func (t *TextAttributes) Fit() model.LeafName { return text.Span }

func (t *TextAttributes) Write(wr model.Writer) {
	wr.Attr(style.Family, style.FamilyText)
	wr.WritePos(New(style.TextProperties))
	if t.fontFace != "" {
		wr.Attr(style.FontName, t.fontFace)
	}
	if t.size != 0 {
		wr.Attr(fo.FontSize, t.size)
	}
	wr.Attr(fo.Color, t.col)
	if t.bold {
		wr.Attr(fo.FontWeight, fo.Bold)
	}
	if t.italic {
		wr.Attr(fo.FontStyle, fo.Italic)
	}
}

func (t *TextAttributes) Size(s int) *TextAttributes {
	t.size = s
	return t
}

func (t *TextAttributes) FontFace(name string) *TextAttributes {
	t.fontFace = name
	return t
}

func (t *TextAttributes) Color(col color.Color) *TextAttributes {
	t.col = col
	return t
}

func (t *TextAttributes) Bold() *TextAttributes {
	t.bold = true
	return t
}

func (t *TextAttributes) Italic() *TextAttributes {
	t.italic = true
	return t
}

func init() {
	New = model.LeafFactory
}
