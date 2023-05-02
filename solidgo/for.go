package solidgo

import "reflect"

func For(collection func() []interface{}, renderer func(index int) Element) Element {
	parent := Window.Document.CreateElement("div")
	currentEffect = func() {
		//TODO for now just rebuild whole thing
	}

	for i := 0; i < reflect.ValueOf(collection()).Len(); i++ {
		parent.AppendChild(renderer(i))
	}
	currentEffect = nil
	return parent
}

type Appendable interface {
	AppendTo(Element)
}

type ForStruct[T any] struct {
	Collection func() []T
	Render     func(T) Element
}

func (forR ForStruct[T]) AppendTo(parent Element) {
	currentEffect = func() {
	}
	for _, item := range forR.Collection() {
		parent.AppendChild(forR.Render(item))
	}
	currentEffect = nil

}

// Generic implementation of For
func ForG[T any](collection func() []T, renderer func(item T) Element) Element {
	parent := Window.Document.CreateElement("div")
	currentEffect = func() {
	}
	for _, item := range collection() {
		parent.AppendChild(renderer(item))
	}
	currentEffect = nil
	return parent
}
