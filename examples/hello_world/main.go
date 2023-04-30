package main

import (
	"github.com/kirillrdy/web/solidgo"
)

type Person struct {
	name string
}

func main() {
	A, T, On := solidgo.A, solidgo.T, solidgo.On
	createEffect := solidgo.CreateEffect
	document := solidgo.Window.Document
	body := document.Body
	getName, setName := solidgo.CreateSignal("Kirill")

	people, setPeople := solidgo.CreateSignal([]Person{
		{name: "Kirill"},
		{name: "Steve"},
		{name: "Bob"},
	})

	selected, setSelected := solidgo.CreateSignal(Person{})

	for _, person := range people() {
		person := person
		div := solidgo.Window.Document.CreateElement("div")
		div.SetInnerText(person.name)
		solidgo.CreateEffect(func() {
			if selected() == person {
				div.SetAttribute("style", "color: red")
			} else {
				div.SetAttribute("style", "")
			}
		})
		div.AddEventListener("click", func() {
			setSelected(person)
		})
		solidgo.Window.Document.Body.AppendChild(div)
	}

	body.AppendChild(solidgo.For(people, func(person Person) solidgo.Element {
		div := document.CreateElement("div")
		div.SetInnerText(person.name)
		createEffect(func() {
			if selected() == person {
				div.SetAttribute("style", "color: red")
			} else {
				div.SetAttribute("style", "")
			}
		})

		return div
	}))

	div := document.CreateElement("div")
	createEffect(func() {
		div.SetInnerText("hello from " + getName())
	})

	body.AppendChild(div)

	div.AddEventListener("click", func() {
		setName("Steve")
		people := append(people(), Person{"New person"})
		setPeople(people)
	})

	body.AppendChild(
		A("div")(On("click", func() { setName("Steve") }))(
			T(func() string { return "Hello world " + getName() }),
		))
	<-make(chan bool)
}
