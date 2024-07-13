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

	var initialPeople []*Person
	for i := 0; i < 10000; i++ {
		initialPeople = append(initialPeople, &Person{name: "Kirill"})
	}

	people, setPeople := solidgo.CreateSignal(initialPeople)

	selected, setSelected := solidgo.CreateSignal(&Person{})

	addPerson := func() {
		people := append(people(), &Person{name: "new guys"})
		setPeople(people)
	}
	personRender := func(person *Person) solidgo.Element {
		style := func() string {
			if selected() == person {
				return "color: red"
			}
			return ""
		}
		return A("div")(At("style", style), On("click", func() { setSelected(person) }))(
			T(person.name),
		)
	}

	A("div")()(
		A("div")()(
			solidgo.For(people, personRender),
		),
		A("button")(On("click", addPerson))(
			T("Click Me"),
		),
	).AppendTo(body)
	<-make(chan bool)
}
