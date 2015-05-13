package model

type AttrName string
type LeafName string

//Attribute of node
type Attribute interface {
	String() string
}

//Leaf is node without children nodes
type Leaf interface {
	Name() LeafName
	Attr(AttrName, ...Attribute) Attribute
	Parent(...Node) Node
}

//Node holds chidlren nodes, Carrier in CRM pattern
type Node interface {
	Leaf
	Child(int) Leaf
	IndexOf(Leaf) int
	NofChild() int
}

//Model holds root node and constructs special riders
type Model interface {
	Root() Node
	NewReader(...Reader) Reader
	NewWriter(...Writer) Writer
}

//Reader is a reading rider in CRM pattern
//Reader stands on node and runs above it's child nodes
type Reader interface {
	InitFrom(Reader)
	Base() Model
	Read() Leaf
	Eol() bool
	Pos(...Leaf) Leaf
}

//Writer is a modifying rider in CRM pattern
//Writer stands on node and modifies it's children and attributes
type Writer interface {
	InitFrom(Writer)
	Base() Model
	Pos(...Leaf) Leaf
	Write(Leaf, ...Leaf)
	WritePos(Leaf, ...Leaf) Leaf
	Attr(AttrName, interface{}) Writer //for fluid interface
	Delete(Leaf)
}
