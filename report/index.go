package report

import (

	c "../config"
	"net/http"
	"text/template"
	"log"
)

var Tmpl = template.Must(template.ParseGlob("Form/*"))

type Employee struct {
    Id    int
    Name  string
    City string
}

func Index(w http.ResponseWriter, r *http.Request) {
    db := c.DbConn()
    selDB, err := db.Query("Select id,name,city from users Order by id ASC")
    if err != nil {
        panic(err.Error())
    }
    emp := Employee{}
    res := []Employee{}
    for selDB.Next() {
        var id int
        var name, city string
        err = selDB.Scan(&id, &name, &city)
        if err != nil {
            panic(err.Error())
        }
        emp.Id = id
        emp.Name = name
        emp.City = city
        res = append(res, emp)
    }
    Tmpl.ExecuteTemplate(w, "Index", res)
    defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
    db := c.DbConn()
    nId := r.URL.Query().Get("id")
    selDB, err := db.Query("SELECT id,name,city FROM users WHERE id=?", nId)
    if err != nil {
        panic(err.Error())
    }
    emp := Employee{}
    for selDB.Next() {
        var id int
        var name, city string
        err = selDB.Scan(&id, &name, &city)
        if err != nil {
            panic(err.Error())
        }
        emp.Id = id
        emp.Name = name
        emp.City = city
    }
    Tmpl.ExecuteTemplate(w, "Show", emp)
    defer db.Close()
}

func Delete(w http.ResponseWriter, r *http.Request) {
    db := c.DbConn()
    emp := r.URL.Query().Get("id")
    delForm, err := db.Prepare("DELETE FROM users WHERE id=?")
    if err != nil {
        panic(err.Error())
    }
    delForm.Exec(emp)
    log.Println("DELETE")
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}