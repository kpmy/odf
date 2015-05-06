package stub

/*
	простая реализация модели документа
*/

import (
	"encoding/xml"
	"odf/model"
	"ypk/assert"
)

type sn struct {
	name     model.LeafName
	attr     map[model.AttrName]model.Attribute
	children []model.Leaf
}

func (n *sn) Attr(name model.AttrName, val ...model.Attribute) model.Attribute {
	assert.For(len(val) <= 1, 20, "only one attribute accepted")
	assert.For(name != "", 21)
	if len(val) == 1 {
		n.attr[name] = val[0]
	}
	return n.attr[name]
}

func (n *sn) Child(i int) model.Leaf {
	assert.For(i < len(n.children), 20)
	return n.children[i]
}

func (n *sn) IndexOf(l model.Leaf) (ret int) {
	ret = -1
	for i := 0; i < len(n.children) && ret == -1; i++ {
		if l == n.children[i] {
			ret = i
		}
	}
	return
}

func (n *sn) NofChild() int {
	return len(n.children)
}

func (n *sn) Name() model.LeafName {
	return n.name
}

func (n *sn) init() {
	n.attr = make(map[model.AttrName]model.Attribute)
	n.children = make([]model.Leaf, 0)
}

func (n *sn) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	start.Name.Local = string(n.name)
	for k, v := range n.attr {
		a, err := v.(xml.MarshalerAttr).MarshalXMLAttr(xml.Name{Local: string(k)})
		assert.For(err == nil, 30, err)
		start.Attr = append(start.Attr, a)
	}
	e.EncodeToken(start)
	for _, _v := range n.children {
		switch v := _v.(type) {
		case *text:
			err = e.EncodeToken(xml.CharData(v.data))
		default:
			err = e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: string(v.Name())}})
		}
		assert.For(err == nil, 30, err)
	}
	err = e.EncodeToken(start.End())
	assert.For(err == nil, 30, err)
	return err
}

type sm struct {
	root *sn
}

func (m *sm) NewReader(old ...model.Reader) model.Reader {
	r := &sr{base: m, eol: true}
	if len(old) == 1 {
		r.InitFrom(old[0])
	}
	return r
}

func (m *sm) NewWriter(old ...model.Writer) model.Writer {
	w := &sw{base: m}
	if len(old) == 1 {
		w.InitFrom(old[0])
	}
	return w
}

func (m *sm) Root() model.Node {
	return m.root
}

func nf() func() model.Model {
	return func() model.Model {
		r := &sn{}
		r.init()
		return &sm{root: r}
	}
}

func init() {
	model.ModelFactory = nf()
}
