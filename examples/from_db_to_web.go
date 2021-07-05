package main

import (
	"github.com/kirillrdy/web/db"
	. "github.com/kirillrdy/web/html"
)

var tables = struct {
	users db.Table
	posts db.Table
}{
	"users",
	"posts",
}
var users = struct {
	id   db.Column
	name db.Column
}{
	tables.users.Column("id"),
	tables.users.Column("name"),
}
var posts = struct {
	title  db.Column
	userId db.Column
}{
	tables.posts.Column("title"),
	tables.posts.Column("user_id"),
}

func main() {

	rows := db.Select(posts.title, users.name).From(tables.posts).Join(tables.users, users.id, posts.userId).Execute()
	for _, row := range rows {
		Div(Class("foo"))(
			P()(
				Text(row.String(posts.title)),
				Text(row.String(users.name)),
			),
		)
	}
}
