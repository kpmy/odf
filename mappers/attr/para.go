package attr

import (
	"github.com/kpmy/odf/model"
	"github.com/kpmy/odf/xmlns/fo"
	"github.com/kpmy/odf/xmlns/style"
	"github.com/kpmy/odf/xmlns/text"
)

//ParagraphAttributes is ODF Paragraph Family style fluent builder
type ParagraphAttributes struct {
	named
	easy
}

func (p *ParagraphAttributes) Equal(_a Attributes) (ok bool) {
	a, ok := _a.(*ParagraphAttributes)
	if ok {
		ok = p.equal(&a.easy)
	}
	return
}

func (p *ParagraphAttributes) Fit() model.LeafName { return text.P }

func (p *ParagraphAttributes) Write(wr model.Writer) {
	wr.Attr(style.Family, style.FamilyParagraph)
	wr.WritePos(New(style.ParagraphProperties))
	p.apply(wr)
}

//AlignRight on page
func (p *ParagraphAttributes) AlignRight() *ParagraphAttributes {
	p.put(fo.TextAlign, fo.Right, nil)
	return p
}

//AlignCenter on page
func (p *ParagraphAttributes) AlignCenter() *ParagraphAttributes {
	p.put(fo.TextAlign, fo.Center, nil)
	return p
}

//PageBreak with new paragraph written (it will be first on new page)
func (p *ParagraphAttributes) PageBreak() *ParagraphAttributes {
	p.put(fo.BreakBefore, true, func(v value) {
		if x := v.data.(bool); x {
			v.wr.Attr(fo.BreakBefore, fo.Page)
		}
	})
	return p
}
