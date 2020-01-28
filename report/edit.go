package report

import (

	c "../config"
	"net/http"
	"log"
)

func Edit(w http.ResponseWriter, r *http.Request) {
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
    Tmpl.ExecuteTemplate(w, "Edit", emp)
    defer db.Close()
}

func Update(w http.ResponseWriter, r *http.Request) {
    db := c.DbConn()
    if r.Method == "POST" {
        name := r.FormValue("name")
        city := r.FormValue("city")
        id := r.FormValue("uid")
        insForm, err := db.Prepare("UPDATE users SET name=?, city=? WHERE id=?")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(name, city, id)
        log.Println("UPDATE: Name: " + name + " | City: " + city)
    }
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}