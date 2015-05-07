package attr

import (
	"odf/model"
	"odf/xmlns/fo"
	"odf/xmlns/style"
	"odf/xmlns/text"
)

type ParagraphAttributes struct {
	align     string
	pageBreak bool
	named
}

func (p *ParagraphAttributes) Equal(_a Attributes) (ok bool) {
	a, ok := _a.(*ParagraphAttributes)
	if ok {
		ok = p.align == a.align && p.pageBreak == a.pageBreak
	}
	return
}

func (p *ParagraphAttributes) Fit() model.LeafName { return text.P }

func (p *ParagraphAttributes) Write(wr model.Writer) {
	wr.Attr(style.Family, style.FamilyParagraph)
	wr.WritePos(New(style.ParagraphProperties))
	wr.Attr(fo.TextAlign, p.align)
	if p.pageBreak {
		wr.Attr(fo.BreakBefore, fo.Page)
	}
}

func (p *ParagraphAttributes) AlignRight() *ParagraphAttributes {
	p.align = fo.Right
	return p
}

func (p *ParagraphAttributes) PageBreak() *ParagraphAttributes {
	p.pageBreak = true
	return p
}
