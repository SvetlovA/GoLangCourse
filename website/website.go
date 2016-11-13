// website
package main

import (
	"html/template"
	"log"
	"net/http"
)

var view = "<html>" +
	"<h1>{{.Title}}</h1>" +
	"<table>" +
	"<tr>" +
	"<th>First name</th>" +
	"<th>Email</th>" +
	"<th>Username</th>" +
	"</tr>" +
	"{{range .Persons}}" +
	"<tr>" +
	"<td>{{.FirstName}}</td>" +
	"<td>{{.Email}}</td>" +
	"<td>{{.Username}}</td>" +
	"</tr>" +
	"{{end}}" +
	"</table>" +
	"</html>"

type Website struct {
	Title   string
	Persons []Person
}

type Person struct {
	FirstName string
	Email     string
	Username  string
}

type Controller struct {
	websit Website
}

func (c *Controller) InitializeTable() {
	c.websit.Title = "User table"
	c.websit.Persons = []Person{
		Person{FirstName: "Artem", Email: "21arte62@gmail.com", Username: "artem"},
		Person{FirstName: "Kirill", Email: "kirill@gmail.com", Username: "kirill"},
		Person{FirstName: "Dmitry", Email: "dmitry@gmail.com", Username: "dmitry"},
	}
}

func (c *Controller) CreateTable(w http.ResponseWriter, r *http.Request) {
	data := template.New("website")
	data, err := data.Parse(view)
	if err != nil {
		log.Fatal(err)
	}

	err = data.Execute(w, c.websit)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var controller Controller

	controller.InitializeTable()
	http.HandleFunc("/", controller.CreateTable)
	http.ListenAndServe(":8080", nil)
}
