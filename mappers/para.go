package mappers

import (
	"odf/model"
	"odf/xmlns/text"
	"ypk/assert"
)

type ParaMapper struct {
	fm    *Formatter
	rider model.Writer
}

func (p *ParaMapper) ConnectTo(fm *Formatter) {
	p.fm = fm
	p.rider = fm.m.NewWriter()
	p.rider.Pos(fm.root)
}

func (p *ParaMapper) WritePara(s string) {
	if pos := p.rider.Pos(); pos.Name() == text.P {
		p.rider.Pos(pos.Parent())
	}
	p.rider.WritePos(New(text.P))
	p.fm.writeAttr()
	if a := p.fm.attr.Fit(text.P); a != nil {
		p.rider.Attr(text.StyleName, a.Name())
	}
	p.WriteString(s)

}
func (p *ParaMapper) WriteString(_s string) {
	assert.For(p.fm.ready, 20)

	buf := make([]rune, 0)
	count := 0

	flush := func(space bool) {
		p.rider.Write(model.Text(string(buf)))
		buf = make([]rune, 0)
		if space && count > 1 {
			w := p.fm.m.NewWriter(p.rider)
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

	if p.rider.Pos().Name() != text.P {
		p.WritePara(_s)
	} else {
		p.fm.writeAttr()
		if a := p.fm.attr.Fit(text.Span); a != nil {
			p.rider.WritePos(New(text.Span))
			p.rider.Attr(text.StyleName, a.Name())
		}
		s := []rune(_s)
		br := false
		for pos := 0; pos < len(s) && s[pos] != 0; {
			switch s[pos] {
			case ' ':
				count++
			case '\n':
				grow()
				p.rider.Write(New(text.LineBreak))
			case '\r':
				grow()
				if p.fm.attr.Fit(text.Span) != nil {
					p.rider.Pos(p.rider.Pos().Parent())
				}
				for pos = pos + 1; pos < len(s); pos++ {
					buf = append(buf, s[pos])
				}
				p.WritePara(string(buf))
				pos--
				br = true
			case '\t':
				grow()
				p.rider.Write(New(text.Tab))
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
			if p.fm.attr.Fit(text.Span) != nil {
				p.rider.Pos(p.rider.Pos().Parent())
			}
		}
	}
}
