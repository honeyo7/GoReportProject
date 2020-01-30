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

	selDB, err :=db.Prepare(strQuery)
	res, err := selDB.Exec()
	var intLastID int64
	if err !=nil{
		intLastID=0
	}else{
		intLastID,err = res.LastInsertId()
		if err !=nil{
			intLastID=0
		}
	}
	
	defer db.Close()
	return intLastID,err

}

func ExecuteUpdateGetRowsAffected(strQuery string) (int64, error){
	db := c.DbConn()

	selDB, err :=db.Prepare(strQuery)
	res, err := selDB.Exec()
	var intAffectedRows int64
	if err !=nil{
		intAffectedRows=0
	}else{
		intAffectedRows,err = res.RowsAffected()
		if err !=nil{
			intAffectedRows=0
		}
	}
	
	defer db.Close()
	return intAffectedRows,err

}

func TestExecuteUpdateGetRowsAffected(strQuery string) (int64, error){
	db := c.DbConn()

	selDB, err :=db.ExecContext(ctx,strQuery)
	var intAffectedRows int64
	if err !=nil{
		intAffectedRows=0
	}else{
		intAffectedRows,err = selDB.RowsAffected()
		if err !=nil{
			intAffectedRows=0
		}
	}
	
	defer db.Close()
	return intAffectedRows,err

}

func TestExecuteInsertGetLastID(strQuery string) (int64, error){
	db := c.DbConn()

	selDB, err :=db.ExecContext(ctx,strQuery)
	var intLastID int64
	if err !=nil{
		intLastID=0
	}else{
		intLastID,err = selDB.LastInsertId()
		if err !=nil{
			intLastID=0
		}
	}
	
	defer db.Close()
	return intLastID,err

}




