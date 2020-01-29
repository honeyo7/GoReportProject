package report

import (

	c "../config"
	"net/http"
    "log"
    odbc "../common"
    "strconv"
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
   
    if r.Method == "POST" {
        name := r.FormValue("name")
        city := r.FormValue("city")
        id := r.FormValue("uid")
        insForm, err := odbc.ExecuteUpdateGetRowsAffected("UPDATE users SET name='" + name + "', city='" + city + "' WHERE id=" + id)
        if err != nil {
            panic(err.Error())
        }

        s1 := strconv.FormatInt(int64(insForm), 10)
        
        log.Println("UPDATE: Name: " + name + " | City: " + city + ". Row Updated:" + s1)
    }
   
    http.Redirect(w, r, "/", 301)
}