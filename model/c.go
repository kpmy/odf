package model

type AttrName string
type LeafName string

var LeafFactory func(LeafName) Leaf
