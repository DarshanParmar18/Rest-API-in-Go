package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDb()  {
	var err error
	DB,err = sql.Open("sqlite3","api.db")//open the database
 
	if err!=nil {
		println(err)
		panic("Could connect to database")
	}

	DB.SetMaxOpenConns(10)	//how many open conn we can have to this database // if the req is more than e.g(10) other req will have to wait until the conn is available again
	DB.SetMaxIdleConns(5)	//how many conn we should keep open if no one is using the conn

	createTables();
}

func createTables() { 		//create the table to store the data into database

	createUsersTable :=`
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`
	_,err:=DB.Exec(createUsersTable)
	if err!=nil {
		panic("cannot create users table")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_,err = DB.Exec(createEventsTable) //TO execute the query

	if err!=nil {
		panic(fmt.Sprintf("could not create events table: %v", err))
	}
}