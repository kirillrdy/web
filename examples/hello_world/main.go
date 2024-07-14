package main

import (
	"github.com/kirillrdy/web/solidgo"
)

type Person struct {
	name string
}

func main() {
	A, T, At, On := solidgo.A, solidgo.T, solidgo.At, solidgo.On

	var initialPeople []*Person
	for i := 0; i < 10000; i++ {
		initialPeople = append(initialPeople, &Person{name: "Kirill"})
	}

	people, _ := solidgo.CreateSignal(initialPeople)

	selected, setSelected := solidgo.CreateSignal(&Person{})

	personRender := func(person *Person) solidgo.Element {
		style := func() string {
			if selected() == person {
				return "color: red"
			}
			return "inherit"
		}
		return A("div")(At("style", style), On("click", func() { setSelected(person) }))(
			T(person.name),
		)
	}

	A("div")()(
		solidgo.For(people, personRender),
	).AppendTo(solidgo.Window.Document.Body)
	<-make(chan bool)
}
