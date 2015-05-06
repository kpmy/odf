package attr

import (
	"odf/model"
	"odf/xmlns/fo"
	"odf/xmlns/style"
	"odf/xmlns/text"
	"ypk/assert"
)

var New func(name model.LeafName) model.Leaf

type Attributes interface {
	Name(...string) string
	Equal(Attributes) bool
	Fit() model.LeafName
	Write(model.Writer)
}

type named struct {
	name string
}

func (n *named) Name(s ...string) string {
	if len(s) == 1 {
		assert.For(s[0] != "", 20)
		n.name = s[0]
	}
	return n.name
}

type TextAttributes struct {
	fontFace string
	size     int
	named
}

func (t *TextAttributes) Equal(_a Attributes) (ok bool) {
	a, ok := _a.(*TextAttributes)
	ok = t.size == a.size && t.fontFace == a.fontFace
	return
}

func (t *TextAttributes) Fit() model.LeafName { return text.Span }

func (t *TextAttributes) Write(wr model.Writer) {
	wr.Attr(style.Family, style.FamilyText)
	wr.WritePos(New(style.TextProperties))
	wr.Attr(style.FontName, t.fontFace)
	wr.Attr(fo.FontSize, t.size)
}

func (t *TextAttributes) Size(s int) *TextAttributes {
	t.size = s
	return t
}

func (t *TextAttributes) FontFace(name string) *TextAttributes {
	t.fontFace = name
	return t
}

func init() {
	New = model.LeafFactory
}
