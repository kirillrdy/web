package main

import (
	"github.com/kirillrdy/web/solidgo"
)

type Person struct {
	name string
}

func main() {
	A, T, At, On := solidgo.A, solidgo.T, solidgo.At, solidgo.On
	body := solidgo.Window.Document.Body

	people, setPeople := solidgo.CreateSignal([]Person{
		{name: "Kirill"},
		{name: "Steve"},
		{name: "Bob"},
	})

	selected, setSelected := solidgo.CreateSignal(Person{})
	addPerson := func() {
		people := append(people(), Person{name: "new guys"})
		setPeople(people)
	}

	A("div")()(
		A("div")()(
			solidgo.For(people, func(person Person) solidgo.Element {
				return A("div")(At("style", func() string {
					if selected() == person {
						return "color: red"
					}
					return ""
				}), On("click", func() { setSelected(person) }))(T(func() string { return person.name }))
			}),
		),
		A("button")(On("click", addPerson))(
			solidgo.Window.Document.CreateTextNode("Click Me"),
		),
	).AppendTo(body)
	<-make(chan bool)
}
