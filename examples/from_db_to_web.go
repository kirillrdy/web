package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path"
	"syscall"
	"time"

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
	id         db.Column
	title      db.Column
	year       db.Column
	created_at db.Column
}{
	tables.movies.Column("id"),
	tables.movies.Column("title"),
	tables.movies.Column("year"),
	tables.movies.Column("created_at"),
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
	postgres := postgresql.Server{DBDir: dbDir}
	if _, err := os.Stat(dbDir); os.IsNotExist(err) == true {
		crash(postgres.InitDB())
		crash(postgres.Start())
		time.Sleep(time.Second)

		crash(postgres.CreateDB(dbName))
		crash(exec.Command("psql", "-d", dbName, "-f", "schema.sql").Run())
		postgres.Stop()
	}

	postgres.Start()
	defer postgres.Stop()

	connection, err := postgres.Connect(dbName)
	crash(err)

	if true {
		for i := 0; i < 1000; i++ {
			db.Insert().Into(tables.movies, movies.title, movies.year).Values("Shawshank Redeption", 1994).Execute(connection)
		}
	}

	movies := admin.Resource{Table: tables.movies}

	admin.AddResource(connection, movies)
	go http.ListenAndServe(":3000", nil)

	ch := make(chan os.Signal, 10)
	signal.Notify(ch, syscall.SIGINT)
	<-ch
}
