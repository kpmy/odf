package mappers

import (
	"github.com/kpmy/ypk/assert"
	"odf/mappers/attr"
	"odf/model"
	"odf/xmlns"
	"odf/xmlns/office"
	"odf/xmlns/text"
	"reflect"
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

func (f *Formatter) makePara() {
	if pos := f.rider.Pos(); pos.Name() != office.Text || pos.Name() == text.P {
		f.rider.Pos(f.text)
	}
	f.rider.WritePos(New(text.P))
	f.attr.Flush()
	f.attr.Fit(text.P, func(a attr.Attributes) {
		f.rider.Attr(text.StyleName, a.Name())
	})
}

func (f *Formatter) WritePara(s string) {
	assert.For(f.ready, 20)
	assert.For(f.MimeType == xmlns.MimeText, 21)
	f.makePara()
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
			w := f.m.NewWriter(f.rider)
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

	if f.rider.Pos().Name() != text.P {
		f.WritePara(_s)
	} else {
		f.attr.Flush()
		f.attr.Fit(text.Span, func(a attr.Attributes) {
			f.rider.WritePos(New(text.Span))
			f.rider.Attr(text.StyleName, a.Name())
		})
		s := []rune(_s)
		br := false
		for pos := 0; pos < len(s) && s[pos] != 0; {
			switch s[pos] {
			case ' ':
				count++
			case '\n':
				grow()
				f.rider.Write(New(text.LineBreak))
			case '\r':
				grow()
				f.attr.Fit(text.Span, func(a attr.Attributes) {
					f.rider.Pos(f.rider.Pos().Parent())
				})
				for pos = pos + 1; pos < len(s); pos++ {
					buf = append(buf, s[pos])
				}
				f.WritePara(string(buf))
				pos--
				br = true
			case '\t':
				grow()
				f.rider.Write(New(text.Tab))
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
			f.attr.Fit(text.Span, func(attr.Attributes) {
				f.rider.Pos(f.rider.Pos().Parent())
			})
		}
	}
}

func (f *Formatter) SetAttr(a attr.Attributes) {
	assert.For(f.ready, 20)
	if a != nil {
		n := reflect.TypeOf(a).String()
		if old := f.attr.OldAttr(a); old != nil {
			f.attr.current[n] = old
		} else {
			c := f.attr.current[n]
			if (c == nil) || !c.Equal(a) {
				f.attr.stored = false
				f.attr.current[n] = a
			}
		}
	} else {
		f.attr.reset()
	}
}

func (f *Formatter) RegisterFont(name, fontface string) {
	f.attr.RegisterFont(name, fontface)
}

func (f *Formatter) SetDefaults(a ...attr.Attributes) {
	f.attr.SetDefaults(a...)
}

func init() {
	New = model.LeafFactory
}
