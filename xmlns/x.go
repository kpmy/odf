package xmlns

import (
	"odf/model"
)

type AttrType int

const (
	NONE AttrType = iota
	STRING
)

type Mime string

const (
	MimeText Mime = "application/vnd.oasis.opendocument.text"
)

var Typed map[model.AttrName]AttrType

func init() {
	Typed = make(map[model.AttrName]AttrType)
}
