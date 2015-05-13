package model

//LeafFactory for pimpl, creates nodes by specified name (names in xmlns/* package and corellate to ODF description
var LeafFactory func(LeafName) Leaf

//ModelFactory for pimpl that creates model
var ModelFactory func() Model

//Text is special leaf constructor because Text Leaf has no name in ODF document model
var Text func(s string) Leaf
