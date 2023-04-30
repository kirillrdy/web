package solidgo

import "syscall/js"

type Document struct {
	Js   js.Value
	Body Element
}
type Element js.Value

var Window = struct {
	Document Document
	Body     Element
}{
	Document: Document{Js: js.Global().Get("document"),
		Body: Element(js.Global().Get("document").Get("body")),
	}}

func (document Document) CreateElement(name string) Element {
	return Element(document.Js.Call("createElement", name))
}

func (document Document) CreateTextNode(content string) Element {
	return Element(document.Js.Call("createTextNode", content))
}

func (element Element) SetAttribute(name, value string) {
	js.Value(element).Call("setAttribute", name, value)
}

func (element Element) AppendChild(value Element) {
	js.Value(element).Call("appendChild", js.Value(value))
}

func (element Element) SetInnerText(value string) {
	js.Value(element).Set("innerText", value)
}

func (element Element) AddEventListener(name string, function func()) {
	callback := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		function()
		return nil
	})
	js.Value(element).Call("addEventListener", name, callback)
}
