package admin

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/kirillrdy/web/db"
	"github.com/kirillrdy/web/html"
	_ "github.com/lib/pq"
)

func page(title string, content html.Node) html.Node {
	Html, Meta, Head, Title, Text, Body := html.Html, html.Meta, html.Head, html.Title, html.Text, html.Body
	Charset := html.Charset
	return Html()(
		Head()(
			Meta(Charset("UTF-8"))(),
			Title()(Text(title)),
			html.Style()(html.TextUnsafe(`
        body {
          font-family: 'Helvetica Neue', Helvetica, Arial, sans-serif;
          line-height: 1.5;
          font-size: 72%;
          color: #323537;
        }
        a {
          margin-right: 7px;
        }
      `)),
		),
		Body()(
			content,
		),
	)
}

func makeIndexHandler(connection *sql.DB, table db.Table, columns []db.Column) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		start := time.Now()
		rows := db.Select(columns...).From(table).Limit(30).
			Execute(connection)
		log.Printf("query took %s", time.Since(start))

		Text, Repeat := html.Text, html.Repeat
		Table, Tr, Th, Div, H1 := html.Table, html.Tr, html.Th, html.Div, html.H1
		Thead, Tbody := html.Thead, html.Tbody

		var tableRows []html.Node
		var header []html.Node
		for _, column := range columns {
			header = append(header, Th()(Text(column.Name)))
		}
		tableRows = append(tableRows, Tr()(header...))

		rowRenderer := func(row db.Row) html.Node {
			var nodes []html.Node
			for _, column := range columns {
				nodes = append(nodes, html.Td()(html.Text(row.String(column))))
			}
			nodes = append(nodes, editLinks(row, table, table.PrimaryKey()))
			return html.Tr()(nodes...)
		}

		page(string(table), Div()(
			H1()(Text(string(table))),
			Table()(
				Thead()(
					tableRows...,
				),
				Tbody()(
					Repeat(rows, rowRenderer)...,
				),
			),
		),
		).WriteTo(response)

	}
}

func PathForRowShow(table db.Table, id int64) string {
	//TODO need to escape
	return fmt.Sprintf("/%s/show?id=%d", table, id)
}

func PathForRowEdit(table db.Table, id int64) string {
	return fmt.Sprintf("/%s/edit?id=%d", table, id)
}

func PathFor(table db.Table) string {
	return "/" + string(table)
}

func editLinks(row db.Row, table db.Table, column db.Column) html.Node {
	A, Href, Text := html.A, html.Href, html.Text
	return html.Td()(
		A(Href(PathForRowShow(table, row.GetInt64(column))))(Text("View")),
		A(Href(PathForRowEdit(table, row.GetInt64(column))))(Text("Edit")),
		A(Href(PathForRowShow(table, row.GetInt64(column))))(Text("Delete")),
	)
}

func makeEditHandler(connection *sql.DB, table db.Table, columns []db.Column) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {

		request.ParseForm()
		id := request.Form["id"][0]
		row, _ := db.Select(columns...).From(table).Where(table.PrimaryKey().Eq(id)).ExecuteOne(connection)

		if request.Method == "POST" {
			log.Printf("%v", request.Form)
			//TODO handle error
			updateQuery := db.Update(table)
			for _, column := range columns {
				if column.PrimaryKey() {
					continue
				}
				updateQuery = updateQuery.Set(column, request.Form[column.FullName()][0])
			}
			updateQuery = updateQuery.Where(table.PrimaryKey().Eq(id))
			// TODO check errors
			updateQuery.Execute(connection)
			http.Redirect(response, request, PathFor(table), 302)
		} else {

			Form, Input, Label, Text := html.Form, html.Input, html.Label, html.Text
			Div, Value, Name, Type, Action, Method := html.Div, html.Value, html.Name, html.Type, html.Action, html.Method
			//Type := html.Type
			var fields []html.Node
			for _, column := range columns {
				if column.PrimaryKey() {
					continue
				}
				input := Div()(
					Label()(Text(column.Name)),
					Input(Value(row.String(column)), Name(column.FullName()))(),
				)
				fields = append(fields, input)
			}
			fields = append(fields, Input(Type("submit"))())
			page(string(table),
				Form(Action(PathForRowEdit(table, row.GetInt64(table.PrimaryKey()))), Method("POST"))(
					fields...,
				),
			).WriteTo(response)
		}
	}
}

type Resource struct {
	Table db.Table
}

func AddResource(connection *sql.DB, resoure Resource) {
	http.HandleFunc("/"+string(resoure.Table)+"/edit", makeEditHandler(connection, resoure.Table, resoure.Table.Columns()))
	http.HandleFunc(PathFor(resoure.Table), makeIndexHandler(connection, resoure.Table, resoure.Table.Columns()))
}
