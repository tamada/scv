package vector

import "fmt"

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

type defaultSource struct {
	kind  string
	value string
}

func NewSource(kind, value string) Source {
	return &defaultSource{kind: kind, value: value}
}

func (ds *defaultSource) Type() string {
	return ds.kind
}

func (ds *defaultSource) Value() string {
	return ds.value
}

type delegateSource struct {
	other Source
}

func (ds *delegateSource) Type() string {
	return fmt.Sprintf("delegate (%s)", ds.other.Type())
}

func (ds *delegateSource) Value() string {
	return ds.other.Value()
}

func delegate(other Source) Source {
	return &delegateSource{other: other}
}
