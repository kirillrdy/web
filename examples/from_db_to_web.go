package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"

	"github.com/kirillrdy/web/admin"
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
	homedir, err := os.UserHomeDir()
	dbDir := path.Join(homedir, "example_db")
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

	if false {
		for i := 0; i < 1000; i++ {
			db.Insert().Into(tables.movies, movies.title).Values("Shawshank Redeption").Execute(connection)
		}
	}

	movies := admin.Resource{Table: tables.movies}

	admin.AddResource(connection, movies)
	http.ListenAndServe(":3000", nil)
}
