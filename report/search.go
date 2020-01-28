
package report

import (

	c "../config"
	"net/http"
	"log"
)


func Search(w http.ResponseWriter, r *http.Request) {
    db := c.DbConn()

    if r.Method == "POST" {
        name := r.FormValue("UName")
        city := r.FormValue("UCity")
        strSearch := ""

        if name!=""{
            strSearch = strSearch + " and name like '%" + name + "%' "
        } 

        if city!=""{
            strSearch = strSearch + " and city like '%" + city + "%' "
        } 

        strQuery := "Select id,name,city from users WHERE id>0 "  + strSearch + " Order by id ASC"

        searchDB, err := db.Query(strQuery)

        if err != nil {
            panic(err.Error())
        }
        emp := Employee{}
        res := []Employee{}
        for searchDB.Next() {
            var id int
            var name, city string
            err = searchDB.Scan(&id, &name, &city)
            if err != nil {
                panic(err.Error())
            }
            emp.Id = id
            emp.Name = name
            emp.City = city
            res = append(res, emp)
        }
        log.Println("Search: Name: " + name + " | City: " + city)
        Tmpl.ExecuteTemplate(w, "Index", res)
        defer db.Close()
    }
    
}