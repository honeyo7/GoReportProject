package report

import (

    "strconv"
	"net/http"
    "log"
    odbc "../odbc"
)

func New(w http.ResponseWriter, r *http.Request) {
    Tmpl.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
   if r.Method == "POST" {
        name := r.FormValue("name")
        city := r.FormValue("city")
        insForm, err := odbc.ExecuteInsertGetLastID("INSERT INTO users(name, city) VALUES('" + name + "','" + city + "')")
        if err != nil {
            panic(err.Error())
        }
        s1 := strconv.FormatInt(int64(insForm), 10)
        log.Println("INSERT: Name: " + name + " | City: " + city + ". UserID: " + s1)
    }
   
    http.Redirect(w, r, "/", 301)
}