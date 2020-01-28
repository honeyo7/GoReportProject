package config

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    
)

func DbConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := "root"
    dbName := "demo"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(127.0.0.1:3306)/"+dbName)
    if err != nil {
        panic(err.Error())
    }
    return db
}