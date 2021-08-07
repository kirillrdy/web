package main

import (
	"log"
	"os/exec"

	"github.com/kirillrdy/web/db"
	"github.com/kirillrdy/web/postgresql"
	_ "github.com/lib/pq"
)

var tables = struct {
	people db.Table
	movies db.Table
}{
	"people",
	"movies",
}
var people = struct {
	id         db.Column
	first_name db.Column
	last_name  db.Column
}{
	tables.people.Column("id"),
	tables.people.Column("first_name"),
	tables.people.Column("last_name"),
}
var movies = struct {
	id    db.Column
	title db.Column
}{
	tables.movies.Column("id"),
	tables.movies.Column("title"),
}

func crash(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	server := postgresql.Server{DBDir: "my_db"}
	server.InitDB()
	server.Start()
	defer server.Stop()
	dbName := "movies"
	crash(server.CreateDB(dbName))
	crash(exec.Command("psql", "-d", dbName, "-f", "schema.sql").Run())

	connection, err := server.Connect(dbName)
	crash(err)

	rows := db.Select(movies.title).From(tables.movies).Execute(connection)
	for _, row := range rows {
		log.Print(row.String(movies.title))
	}
}
