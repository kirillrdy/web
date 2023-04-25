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

func (element Element) AddEventListener(name string, function func()) {
	callback := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		function()
		return nil
	})
	js.Value(element).Call("addEventListener", name, callback)
}

type Signal[T any] struct {
	storage  T
	notifies []func()
}

func createSignal[T any](defaultValue T) (func() T, func(T)) {
	signal := Signal[T]{storage: defaultValue}
	return signal.Get, signal.Set
}

// TODO panic when called outside effect
func (signal *Signal[T]) Get() T {
	if currectEffect != nil {
		signal.notifies = append(signal.notifies, currectEffect)
	}
	return signal.storage
}
func (signal *Signal[T]) Set(value T) {
	signal.storage = value
	for _, function := range signal.notifies {
		function()
	}
}

var currectEffect func()

func createEffect(function func()) {
	currectEffect = function
	function()
	currectEffect = nil
}

type Person struct {
	name string
}

var document = Document(js.Global().Get("document"))
var body = Element(js.Value(document).Get("body"))

func For[T any](parent Element, collection func() []T, renderer func(item T) Element) {
	currectEffect = func() {
	}
	for _, item := range collection() {
		parent.AppendChild(renderer(item))
	}
	currectEffect = nil
}

func main() {
	getName, setName := createSignal("Kirill")

	people, setPeople := createSignal([]Person{
		{name: "Kirill"},
		{name: "Steve"},
		{name: "Bob"},
	})

	selected, setSelected := createSignal(Person{})

	for _, person := range people() {
		person := person
		div := document.createElement("div")
		div.SetInnerText(person.name)
		createEffect(func() {
			if selected() == person {
				div.SetAttribute("style", "color: red")
			} else {
				div.SetAttribute("style", "")
			}
		})
		div.AddEventListener("click", func() {
			setSelected(person)
		})
		body.AppendChild(div)
	}

	For(body, people, func(person Person) Element {
		div := document.createElement("div")
		div.SetInnerText(person.name)
		createEffect(func() {
			if selected() == person {
				div.SetAttribute("style", "color: red")
			} else {
				div.SetAttribute("style", "")
			}
		})

		return div
	})

	div := document.createElement("div")
	createEffect(func() {
		div.SetInnerText("hello from " + getName())
	})

	body.AppendChild(div)

	div.AddEventListener("click", func() {
		setName("Steve")
		people := append(people(), &Person{"New person"})
		setPeople(people)
	})

	<-make(chan bool)
}
