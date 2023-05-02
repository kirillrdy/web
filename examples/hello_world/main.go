package main

import (
	"github.com/kirillrdy/web/solidgo"
)

type Person struct {
	name string
}

func main() {
	A, At, T, On := solidgo.A, solidgo.At, solidgo.T, solidgo.On
	document := solidgo.Window.Document
	body := document.Body

	people, setPeople := solidgo.CreateSignal([]Person{
		{name: "Kirill"},
		{name: "Steve"},
		{name: "Bob"},
	})

	selected, setSelected := solidgo.CreateSignal(Person{})

	A("div")()(
		A("div")()(
			solidgo.ForStruct[Person]{Collection: people, Render: func(person Person) solidgo.Element {
				return A("div")(At("style", func() string {
					if selected() == person {
						return "color: red"
					}
					return ""
				}), On("click", func() { setSelected(person) }))(T(func() string { return person.name }))
			}}),
		A("button")(On("click", func() {
			people := append(people(), Person{name: "new guys"})
			setPeople(people)
		}))(T(func() string { return "Click Me" })),
	).AppendTo(body)
	<-make(chan bool)
}
