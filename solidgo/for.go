package solidgo

import (
	"reflect"
	"syscall/js"
)

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

type ForStruct[T comparable] struct {
	Collection func() []T
	Render     func(T) Element
}

func (forR ForStruct[T]) AppendTo(parent Element) {
	var previous []Element
	currentEffect = func() {
		for _, item := range previous {
			// replaced Remove with binding
			js.Value(item).Call("remove")
		}
		previous = nil
		for _, item := range forR.Collection() {
			element := forR.Render(item)
			previous = append(previous, element)
			parent.AppendChild(element)
		}
	}
	currentEffect()

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
