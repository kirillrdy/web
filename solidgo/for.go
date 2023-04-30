package solidgo

func For[T any](collection func() []T, renderer func(item T) Element) Element {
	parent := document.createElement("div")
	currentEffect = func() {
	}
	for _, item := range collection() {
		parent.AppendChild(renderer(item))
	}
	currentEffect = nil
	return parent
}
