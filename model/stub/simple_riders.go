package stub

import (
	"odf/model"
	"odf/xmlns"
	"reflect"
	"ypk/assert"
	"ypk/halt"
)

type sw struct {
	base *sm
	pos  model.Leaf
}

func (w *sw) Base() model.Model {
	return w.base
}

func (w *sw) InitFrom(old model.Writer) {
	panic(126)
}

func (w *sw) Pos(p ...model.Leaf) model.Leaf {
	if len(p) == 1 {
		w.pos = p[0]
	}
	return w.pos
}

func (w *sw) Write(l model.Leaf) {
	assert.For(l != nil, 20)
	assert.For(w.pos != nil, 21)
	if _n, ok := w.pos.(model.Node); ok {
		if n, da := _n.(*sn); da {
			n.children = append(n.children, l)
		} else {
			halt.As(100, reflect.TypeOf(n))
		}
	}
}

func (w *sw) WritePos(l model.Leaf) model.Leaf {
	w.Write(l)
	return w.Pos(l)
}

func castAttr(n model.AttrName, i interface{}) (ret model.Attribute) {
	typ := xmlns.Typed[n]
	switch typ {
	case xmlns.NONE, xmlns.STRING:
		ret = &StringAttr{Value: i.(string)}
	default:
		halt.As(100, typ, reflect.TypeOf(i))
	}
	return ret
}

func (w *sw) Attr(n model.AttrName, val interface{}) {
	w.pos.Attr(n, castAttr(n, val))
}
