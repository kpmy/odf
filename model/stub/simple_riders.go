package stub

import (
	"odf/model"
	"odf/xmlns"
	"reflect"
	"ypk/assert"
	"ypk/halt"
)

type sr struct {
	base *sm
	pos  model.Leaf
	eol  bool
	this model.Leaf
}

func (r *sr) Base() model.Model {
	return r.base
}

func (r *sr) InitFrom(old model.Reader) {
	panic(126)
}

func (r *sr) Pos(p ...model.Leaf) model.Leaf {
	if len(p) == 1 {
		r.pos = p[0]
		if n, ok := r.pos.(model.Node); ok {
			r.eol = n.NofChild() == 0
		} else {
			r.eol = true
		}
	}
	return r.pos
}

func (r *sr) Read() model.Leaf {
	assert.For(r.pos != nil && !r.eol, 20)
	n, ok := r.pos.(model.Node)
	assert.For(ok, 21)
	idx := 0
	if r.this != nil {
		idx = n.IndexOf(r.this)
		idx++
	}
	if idx < n.NofChild() {
		r.this = n.Child(idx)
	} else {
		r.eol = true
	}
	return r.this
}

func (r *sr) Eol() bool {
	return r.eol
}

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
		switch n := _n.(type) {
		case *sn:
			n.children = append(n.children, l)
		case *root:
			n.inner.children = append(n.inner.children, l)
		default:
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
