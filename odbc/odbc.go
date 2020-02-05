package odbc

import (

	c "../config"
	"database/sql"
	"context"
	"errors"
	//"database/sql/driver"  Result Interface
)

var (
	ctx context.Context
)

func ExecuteNonQuery(strQuery string) (error){
	db := c.DbConn()
	_, err := db.Query(strQuery)

	defer db.Close()
	return err

	

}

func ExecuteQueryRows(strQuery string) ( *sql.Rows, error){
	db := c.DbConn()
	selDB, err := db.Query(strQuery)
	defer db.Close()
	return selDB,err

}

func ExecuteQueryInt(strQuery string) ( int64, error){
	selDB, err := ExecuteQueryRows(strQuery)
	
	if err != nil {
		return 0,err
	}
	var intResult int64
	count := 0
	for selDB.Next() {
		err = selDB.Scan(&intResult)
        if err != nil {
            panic(err.Error())
		}
		
		count += 1 
	}

	if count==0 {
		err:=errors.New("No row Found!")
		return 0,err
	}

	if count>1 {
		err:=errors.New("Multiple rows Found!")
		return 0,err
	}

	return intResult,err

}

func ExecuteQueryStr(strQuery string) ( string, error){
	selDB, err := ExecuteQueryRows(strQuery)
	
	if err != nil {
		return "",err
	}
	var strResult string
	count := 0
	for selDB.Next() {
		err = selDB.Scan(&strResult)
        if err != nil {
            panic(err.Error())
		}
		
		count += 1 
	}

	if count==0 {
		err:=errors.New("No row Found!")
		return "",err
	}

	if count>1 {
		err:=errors.New("Multiple rows Found!")
		return "",err
	}

	return strResult,err

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