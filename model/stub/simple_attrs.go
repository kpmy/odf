package stub

type StringAttr struct {
	Value string
}

func (a *StringAttr) String() string {
	return a.Value
}
