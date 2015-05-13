package stub

import (
	"encoding/xml"
	"fmt"
	"image/color"
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

type BoolAttr struct {
	Value bool
}

func (a *BoolAttr) String() string {
	if a.Value {
		return "true"
	} else {
		return "false"
	}
}

func (a *BoolAttr) MarshalXMLAttr(name xml.Name) (xa xml.Attr, err error) {
	xa.Name = name
	xa.Value = a.String()
	return
}

type MeasureAttr struct {
	Value float64
}

func (a *MeasureAttr) String() string {
	return strconv.FormatFloat(a.Value, 'f', 8, 64) + "cm"
}

func (a *MeasureAttr) MarshalXMLAttr(name xml.Name) (xa xml.Attr, err error) {
	xa.Name = name
	xa.Value = a.String()
	return
}

type ColorAttr struct {
	Value color.Color
}

func (a *ColorAttr) String() string {
	r, g, b, _ := a.Value.RGBA()
	return fmt.Sprintf("#%02X%02X%02X", uint8(r), uint8(g), uint8(b))
}

func (a *ColorAttr) MarshalXMLAttr(name xml.Name) (xa xml.Attr, err error) {
	xa.Name = name
	xa.Value = a.String()
	return
}
