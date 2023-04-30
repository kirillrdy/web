package solidgo

import "syscall/js"

type Document js.Value
type Element js.Value

var document = Document(js.Global().Get("document"))
var body = Element(js.Value(document).Get("body"))

func (document Document) createElement(name string) Element {
	return Element(js.Value(document).Call("createElement", name))
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
