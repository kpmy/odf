package stub

import (
	"odf/model"
	"ypk/assert"
)

func lf() func(x model.LeafName) model.Leaf {
	return func(x model.LeafName) model.Leaf {
		var ret *sn
		switch x {
		default:
			ret = &sn{name: x}
			ret.init()
		}
		assert.For(ret != nil, 60)
		return ret
	}
}

func init() {
	model.LeafFactory = lf()
}
