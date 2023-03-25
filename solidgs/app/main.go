package main

import (
	"syscall/js"
)

type Document js.Value
type Element js.Value

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

var document = Document(js.Global().Get("document"))
var body = Element(js.Value(document).Get("body"))

func main() {
	js.Global().Get("console").Call("log", "hello world")
	div := document.createElement("div")
	div.SetInnerText("hello from kirill")
	body.AppendChild(div)
}
