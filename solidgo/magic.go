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
	//TODO replace this with automatic wrapper
	// eg runEffect(func(){
	// })
	// where runEffect will also execute effect
	currentEffect = func() {
		element.SetAttribute(attribute.Name, attribute.Effect())
	}
	currentEffect()

	currentEffect = nil
}

func On(event string, callback func()) EventCallBack {
	return EventCallBack{Name: event, Callback: callback}
}

func T(effect func() string) Element {
	element := Window.Document.CreateTextNode("")
	currentEffect = func() {
		js.Value(element).Set("textContent", effect())
	}
	currentEffect()
	currentEffect = nil
	return element
}

func At(name string, effect func() string) Attribute {
	return Attribute{Name: name, Effect: effect}
}

func A(name string) func(...ApplyToElement) func(...Appendable) Element {
	return func(attributes ...ApplyToElement) func(...Appendable) Element {
		element := Window.Document.CreateElement(name)
		for _, attribute := range attributes {
			attribute.Apply(element)
		}
		return func(children ...Appendable) Element {
			for _, child := range children {
				child.AppendTo(element)
			}
			return element
		}
	}
}
