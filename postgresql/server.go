package postgresql

import (
	"database/sql"
	"os/exec"
)

type Server struct {
	DBDir string
}

func (server Server) InitDB() error {
	return exec.Command("initdb", server.DBDir).Run()
}

func (server Server) Start() error {
	return exec.Command("pg_ctl", "-D", server.DBDir, "start").Run()
}

func (server Server) Stop() error {
	return exec.Command("pg_ctl", "-D", server.DBDir, "stop").Run()
}

func (server Server) CreateDB(dbName string) error {
	return exec.Command("createdb", dbName).Run()
}
func (server Server) Connect(dbName string) (*sql.DB, error) {
	return sql.Open("postgres", "postgres://localhost/"+dbName+"?sslmode=disable")
}
