package stub

import (
	"encoding/xml"
)

type StringAttr struct {
	Value string
}

func (a *StringAttr) String() string {
	return a.Value
}

func (a *StringAttr) MarshalXMLAttr(name xml.Name) (xa xml.Attr, err error) {
	xa.Name = name
	xa.Value = a.Value
	return
}
