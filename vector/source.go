package vector

type Source interface {
	Type() string
	Value() string
}

type stringSource struct {
	typeValue string
	value     string
}

func (ss *stringSource) Type() string {
	return ss.typeValue
}

func (ss *stringSource) Value() string {
	return ss.value
}

func newStringSource(str string) Source {
	return &stringSource{typeValue: "string", value: str}
}
