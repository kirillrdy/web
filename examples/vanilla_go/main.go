package main

import (
	"github.com/kirillrdy/web/solidgo"
)

func main() {
	A, T := solidgo.A, solidgo.T
	body := solidgo.Window.Document.Body

	var selected *solidgo.Element

	for i := 0; i < 10000; i++ {
		span := A("div")()(T("Kirill"))
		span.AddEventListener("click", func() {
			if selected != nil {
				selected.SetAttribute("style", "color: inherit;")
			}
			span.SetAttribute("style", "color: red;")
			selected = &span
		})
		body.AppendChild(span)
	}
	<-make(chan bool)
}
