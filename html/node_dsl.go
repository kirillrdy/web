package html

import (
	"github.com/kirillrdy/web/db"
)

type ifWrapper struct {
	condition bool
	thenNodes []Node
	elseNodes []Node
}

func Repeat(rows []db.Row, render func(db.Row) Node) []Node {
	var result []Node
	for _, row := range rows {
		result = append(result, render(row))
	}
	return result
}

func If(condition bool) ifWrapper {
	return ifWrapper{condition: condition}
}

func (wrapper ifWrapper) Then(nodes ...Node) ifWrapper {
	wrapper.thenNodes = nodes
	return wrapper
}

func (wrapper ifWrapper) Else(nodes ...Node) ifWrapper {
	wrapper.elseNodes = nodes
	return wrapper
}

func (wrapper ifWrapper) Nodes() []Node {
	if wrapper.condition {
		return wrapper.thenNodes
	} else {
		return wrapper.elseNodes
	}
}
