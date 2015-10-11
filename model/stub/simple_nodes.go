package stub

import (
	"encoding/xml"
	"github.com/kpmy/odf/model"
	"github.com/kpmy/odf/xmlns"
	"github.com/kpmy/odf/xmlns/office"
	"github.com/kpmy/odf/xmlns/urn"
	"github.com/kpmy/ypk/assert"
)

type root struct {
	inner *sn
}

func (r *root) Name() model.LeafName {
	return r.inner.Name()
}

func (r *root) Attr(n model.AttrName, a ...model.Attribute) model.Attribute {
	return r.inner.Attr(n, a...)
}

func (r *root) Child(i int) model.Leaf {
	return r.inner.Child(i)
}

func (r *root) IndexOf(l model.Leaf) int {
	return r.inner.IndexOf(l)
}

func (r *root) NofChild() int {
	return r.inner.NofChild()
}

func (r *root) Parent(...model.Node) model.Node { return nil }

type text struct {
	data   string
	parent model.Node
}

func (t *text) Name() model.LeafName { panic(100) }

func (t *text) Attr(model.AttrName, ...model.Attribute) model.Attribute { panic(100) }

func (t *text) Parent(p ...model.Node) model.Node {
	if len(p) == 1 {
		assert.For(t.parent == nil, 20)
		t.parent = p[0]
	}
	return t.parent
}

func (r *root) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	start.Name.Local = string(r.inner.Name())

	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: xmlns.NSanim}, Value: urn.Anim})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: xmlns.NSchart}, Value: urn.Chart})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: xmlns.NSconfig}, Value: urn.Config})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: xmlns.NSdc}, Value: urn.Dc})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: xmlns.NSdr3d}, Value: urn.Dr3d})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: xmlns.NSdraw}, Value: urn.Draw})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: xmlns.NSfo}, Value: urn.Fo})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: xmlns.NSform}, Value: urn.Form})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: xmlns.NSmath}, Value: urn.Math})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: xmlns.NSmeta}, Value: urn.Meta})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: xmlns.NSnumber}, Value: urn.Number})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: xmlns.NSoffice}, Value: urn.Office})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: xmlns.NSpresentation}, Value: urn.Presentation})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: xmlns.NSscript}, Value: urn.Script})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: xmlns.NSsmil}, Value: urn.Smil})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: xmlns.NSstyle}, Value: urn.Style})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: xmlns.NSsvg}, Value: urn.Svg})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: xmlns.NStable}, Value: urn.Table})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: xmlns.NStext}, Value: urn.Text})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: xmlns.NSxforms}, Value: urn.Xforms})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: xmlns.NSxlink}, Value: urn.Xlink})
	err = e.EncodeElement(r.inner, start)
	return err
}

func lf() func(x model.LeafName) model.Leaf {
	return func(x model.LeafName) (ret model.Leaf) {
		switch x {
		case office.DocumentContent, office.DocumentMeta, office.DocumentStyles:
			r := &root{}
			r.inner = &sn{name: x}
			r.inner.init()
			ret = r
		default:
			r := &sn{name: x}
			r.init()
			ret = r
		}
		assert.For(ret != nil, 60)
		return ret
	}
}

func tf() func(s string) model.Leaf {
	return func(s string) model.Leaf {
		return &text{data: s}
	}
}

func init() {
	model.LeafFactory = lf()
	model.Text = tf()
}
