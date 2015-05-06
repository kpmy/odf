package model

type Attribute interface {
	String() string
}

type Leaf interface {
	Name() LeafName
	Attr(AttrName, ...Attribute) Attribute
}

type Node interface {
	Leaf
	Child(int) Leaf
	IndexOf(Leaf) int
	NofChild() int
}

type Model interface {
	Root() Node
	NewReader(...Reader) Reader
	NewWriter(...Writer) Writer
}

type Reader interface {
	InitFrom(Reader)
	Base() Model
	Read() Leaf
	Eol() bool
	Pos(...Leaf) Leaf
}

type Writer interface {
	InitFrom(Writer)
	Base() Model
	Pos(...Leaf) Leaf
	Write(Leaf)
	WritePos(Leaf) Leaf
	Attr(AttrName, interface{})
}

var ModelFactory func() Model
var Text func(s string) Leaf
