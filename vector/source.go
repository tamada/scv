package vector

type Source interface {
	Type() string
	Value() string
}

type stringSource struct {
	value string
}

func (ss *stringSource) Type() string {
	return "string"
}

func (ss *stringSource) Value() string {
	return ss.value
}

func newStringSource(str string) Source {
	return &stringSource{value: str}
}
