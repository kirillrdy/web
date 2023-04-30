package solidgo

import "syscall/js"

type ApplyToElement interface {
	Apply(Element)
}

type Attribute struct {
	Name   string
	Effect func() string
}

type EventCallBack struct {
	Name     string
	Callback func()
}

func (event EventCallBack) Apply(element Element) {
	element.AddEventListener(event.Name, event.Callback)
}

func (attribute Attribute) Apply(element Element) {
	element.SetAttribute(attribute.Name, attribute.Effect())
}

func On(event string, callback func()) EventCallBack {
	return EventCallBack{Name: event, Callback: callback}
}

func T(effect func() string) Element {
	element := Window.Document.CreateTextNode("")
	currentEffect = func() {
		value := effect()
		js.Value(element).Set("textContent", value)
	}
	currentEffect()
	return element
}

func At(name string, effect func() string) Attribute {
	return Attribute{Name: name, Effect: effect}
}

func A(name string) func(...ApplyToElement) func(...Element) Element {
	element := Window.Document.CreateElement(name)
	attributesFunction := func(attributes ...ApplyToElement) func(...Element) Element {
		for _, attribute := range attributes {
			attribute.Apply(element)
		}
		return func(children ...Element) Element {
			for _, child := range children {
				element.AppendChild(child)
			}
			return element
		}
	}
	return attributesFunction
}
