package vector

type Source interface {
	Type() string
	Value() string
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
