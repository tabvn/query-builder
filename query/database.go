package query

import "database/sql"

type Database struct {
	DB *sql.DB
}

func NewDatabase(url string) *Database{

	db, err := sql.Open("mysql", url)

	if err != nil{
		panic(err)
	}

	return &Database{
		DB: db,
	}
}

func (d *Database) Insert(table string) *InsertQuery{

	return &InsertQuery{db: d, table: table}
}

func (d *Database) Exec(query string, args ...interface{}) (sql.Result, error){
	return d.DB.Exec(query, args...)
}