//Package stub is an implementation of document model, nodes, attributes and riders, it's a lightweight analog of DOM
//There is no need to create separate type for each ODF document node, because most of them are defined only by it's name, that's why we build a simle tree model where only root nodel of document are special, because they are translated into files of document-package. But they still support Node interface and may be used in other implementations as regular nodes
package stub

/*
	простая реализация модели документа
*/

import (
	"encoding/xml"
	"github.com/kpmy/ypk/assert"
	"odf/model"
)

type sn struct {
	name     model.LeafName
	attr     map[model.AttrName]model.Attribute
	children []model.Leaf
	parent   model.Node
}

func (n *sn) Attr(name model.AttrName, val ...model.Attribute) model.Attribute {
	assert.For(len(val) <= 1, 20, "only one attribute accepted")
	assert.For(name != "", 21)
	if len(val) == 1 {
		if val[0] != nil {
			n.attr[name] = val[0]
		} else {
			delete(n.attr, name)
		}
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

func (n *sn) Parent(p ...model.Node) model.Node {
	if len(p) == 1 {
		assert.For(n.parent == nil, 20)
		n.parent = p[0]
	}
	return n.parent
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
