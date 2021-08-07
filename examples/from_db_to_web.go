package main

import (
	"github.com/kirillrdy/web/db"
	. "github.com/kirillrdy/web/html"
	"github.com/kirillrdy/web/postgresql"
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
	server := postgresql.Server{DBDir: "my_db"}
	server.InitDB()
	server.Start()
	defer server.Stop()
	dbName := "movies"
	server.CreateDB(dbName)

	connection, _ := server.Connect(dbName)
	db.DB = connection

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
