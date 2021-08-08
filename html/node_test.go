package html

import "testing"

func BenchmarkNodeToString(b *testing.B) {
	node := Div()
	for i := 0; i < 1000; i++ {
		nested := Div()()
		var nodes []Node
		for i := 0; i < 1000; i++ {
			nodes = append(nodes, Span()(Text("Some text")))
		}
		node(nested)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		node().String()
	}
}
