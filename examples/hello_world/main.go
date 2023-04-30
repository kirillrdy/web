package main

import (
	"github.com/kirillrdy/web/solidgo"
)

type Person struct {
	name string
}

func main() {
	getName, setName := solidgo.CreateSignal("Kirill")

	people, setPeople := solidgo.CreateSignal([]Person{
		{name: "Kirill"},
		{name: "Steve"},
		{name: "Bob"},
	})

	selected, setSelected := solidgo.CreateSignal(Person{})

	for _, person := range people() {
		person := person
		div := solidgo.document.createElement("div")
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

	body.AppendChild(For(people, func(person Person) Element {
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
	}))

	div := document.createElement("div")
	createEffect(func() {
		div.SetInnerText("hello from " + getName())
	})

	body.AppendChild(div)

	div.AddEventListener("click", func() {
		setName("Steve")
		people := append(people(), Person{"New person"})
		setPeople(people)
	})

	<-make(chan bool)
}
