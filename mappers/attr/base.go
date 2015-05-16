package attr

import (
	"github.com/kpmy/ypk/assert"
	"odf/model"
)

//New is shortcut for factory function model.ModelFactory
var New func(name model.LeafName) model.Leaf

//Attributes interface that is supported by current version of mappers.Attr container
type Attributes interface {
	Name(...string) string
	Equal(Attributes) bool
	Fit() model.LeafName
	Write(model.Writer)
}

type named struct {
	name string
}

func (n *named) Name(s ...string) string {
	if len(s) == 1 {
		assert.For(s[0] != "", 20)
		n.name = s[0]
	}
	return n.name
}

type value struct {
	name    model.AttrName
	wr      model.Writer
	data    interface{}
	builder func(value)
}

type easy struct {
	m map[model.AttrName]value
}

func (e *easy) put(n model.AttrName, x interface{}, foo func(value)) {
	if e.m == nil {
		e.m = make(map[model.AttrName]value)
	}
	b := func(v value) {
		v.wr.Attr(v.name, v.data)
	}
	if foo != nil {
		b = foo
	}
	if x != nil {
		e.m[n] = value{data: x, builder: b}
	} else {
		delete(e.m, n)
	}
}

func (e *easy) equal(t *easy) (ok bool) {
	ok = (e.m != nil) == (t.m != nil)
	if ok && (e.m != nil) {
		for k, v := range e.m {
			ok = t.m[k].data == v.data
			if !ok {
				break
			}
		}
	}
	return
}

func (e *easy) apply(wr model.Writer) {
	if e.m != nil {
		for k, v := range e.m {
			v.wr = wr
			v.name = k
			v.builder(v)
		}
	}
}

func triggerBoolAttr(n model.AttrName) func(v value) {
	return func(v value) {
		if x := v.data.(bool); x {
			v.wr.Attr(n, true)
		}
	}
}

func init() {
	New = func(n model.LeafName) model.Leaf {
		return model.LeafFactory(n)
	}
}
