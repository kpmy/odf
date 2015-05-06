package stub

import (
	"encoding/xml"
	"strconv"
)

type StringAttr struct {
	Value string
}

func (a *StringAttr) String() string {
	return a.Value
}

func (a *StringAttr) MarshalXMLAttr(name xml.Name) (xa xml.Attr, err error) {
	xa.Name = name
	xa.Value = a.String()
	return
}

type IntAttr struct {
	Value int
}

func (a *IntAttr) String() string {
	return strconv.Itoa(a.Value)
}

func (a *IntAttr) MarshalXMLAttr(name xml.Name) (xa xml.Attr, err error) {
	xa.Name = name
	xa.Value = a.String()
	return
}
