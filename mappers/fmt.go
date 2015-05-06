package mappers

import (
	"odf/model"
	"odf/xmlns"
	"odf/xmlns/office"
	"odf/xmlns/text"
	"ypk/assert"
)

var New func(name model.LeafName) model.Leaf

type Formatter struct {
	m        model.Model
	rider    model.Writer
	MimeType xmlns.Mime
	attr     *Attr
	text     model.Node
	ready    bool
}

func (f *Formatter) ConnectTo(m model.Model) {
	assert.For(m.Root().NofChild() == 0, 20, "only new documents for now")
	f.m = m
	f.rider = f.m.NewWriter()
	f.attr = &Attr{}
	f.ready = false
}

func (f *Formatter) Init() {
	assert.For(f.MimeType != "", 20)
	wr := f.m.NewWriter()
	wr.Pos(f.m.Root())
	wr.Write(New(office.DocumentMeta))
	f.attr.Init(f.m)
	wr.WritePos(New(office.DocumentContent))
	wr.Attr(office.Version, "1.0")
	wr.Write(f.attr.ffdc)
	wr.Write(f.attr.asc)
	wr.WritePos(New(office.Body))
	f.text = wr.WritePos(New(office.Text)).(model.Node)
	f.rider.Pos(f.text)
	f.ready = true
}

func (f *Formatter) writeAttr() {

}

func (f *Formatter) WritePara(s string) {
	assert.For(f.ready, 20)
	assert.For(f.MimeType == xmlns.MimeText, 21)
	if pos := f.rider.Pos(); pos.Name() != office.Text || pos.Name() == text.Paragraph {
		f.rider.Pos(f.text)
	}
	f.rider.WritePos(New(text.Paragraph))
	f.writeAttr()
	f.WriteString(s)
}

func (f *Formatter) WriteString(_s string) {
	assert.For(f.ready, 20)

	buf := make([]rune, 0)
	count := 0

	flush := func(space bool) {
		f.rider.Write(model.Text(string(buf)))
		buf = make([]rune, 0)
		if space && count > 1 {
			w := f.m.NewWriter()
			w.WritePos(New(text.S))
			w.Attr(text.C, count)
		}
	}

	grow := func() {
		if count > 1 {
			flush(true)
		} else if count == 1 {
			buf = append(buf, ' ')
		}
		if len(buf) > 0 {
			flush(false)
		}
		count = 0
	}

	if f.rider.Pos().Name() != text.Paragraph {
		f.WritePara(_s)
	} else {
		f.writeAttr()
		s := []rune(_s)
		br := false
		for pos := 0; pos < len(s) && s[pos] != 0; {
			switch s[pos] {
			case ' ':
				count++
			default:
				if count > 1 {
					flush(true)
				} else if count == 1 {
					buf = append(buf, ' ')
				}
				count = 0
				buf = append(buf, s[pos])
			}
			pos++
		}
		if !br {
			grow()
		}
	}
}

func init() {
	New = model.LeafFactory
}
