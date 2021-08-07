package main

import (
	"log"
	"os"
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
	dbDir := "example_db"
	dbName := "movies"
	server := postgresql.Server{DBDir: dbDir}
	if _, err := os.Stat(dbDir); os.IsNotExist(err) == true {
		server.InitDB()
		server.Start()

		crash(server.CreateDB(dbName))
		crash(exec.Command("psql", "-d", dbName, "-f", "schema.sql").Run())
		server.Stop()
	}

	server.Start()
	defer server.Stop()

	connection, err := server.Connect(dbName)
	crash(err)

	db.Insert().Into(tables.movies, movies.title).Values("Shawshank Redeption").Execute(connection)

	rows := db.Select(movies.title).From(tables.movies).Execute(connection)
	for _, row := range rows {
		log.Print(row.String(movies.title))
	}
}
