package attr

import (
	"github.com/kpmy/ypk/assert"
	"odf/model"
)

var New func(name model.LeafName) model.Leaf

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
