package odbc

import (

	c "../config"
	"database/sql"
	"context"
	//"database/sql/driver"  Result Interface
)

var (
	ctx context.Context
)

func ExecuteNonQuery(strQuery string) (bool,error){
	db := c.DbConn()
	_, err := db.Query(strQuery)

	var boolResult bool

	if err !=nil{
		boolResult=false
	}else{
		boolResult=true
	}
	defer db.Close()
	return boolResult,err

	

}

func ExecuteQueryRows(strQuery string) ( *sql.Rows, error){
	db := c.DbConn()
	selDB, err := db.Query(strQuery)
	defer db.Close()
	return selDB,err

}

func ExecuteInsertGetLastID(strQuery string) (int64, error){
	db := c.DbConn()
	var intLastID int64
	
	selDB, err :=db.Prepare(strQuery)
	if err !=nil{
		intLastID=0
		return intLastID,err 
	}

	res, err := selDB.Exec()
	if err !=nil{
		intLastID=0
		return intLastID,err 
	}
	
	
	intLastID,err = res.LastInsertId()
	if err !=nil{
		intLastID=0
	}
	
	
	defer db.Close()
	return intLastID,err

}

func ExecuteUpdateGetRowsAffected(strQuery string) (int64, error){
	db := c.DbConn()
	var intAffectedRows int64
	selDB, err :=db.Prepare(strQuery)
	if err !=nil{
		intAffectedRows=0
		return intAffectedRows,err 
	}
	
	res, err := selDB.Exec()
	if err !=nil{
		intAffectedRows=0
		return intAffectedRows,err 
	}
	
	intAffectedRows,err = res.RowsAffected()
	if err !=nil{
		intAffectedRows=0
	}
	
	
	defer db.Close()
	return intAffectedRows,err

}




