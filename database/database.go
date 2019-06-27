package database

import(
	"database/sql"
	"log"
	"os"
	_"github.com/lib/pq"
)

type database struct{
	url string
}

func (db database) InitialDB(){
	url = os.Getenv("DATABASE_URL")
	if len(url) == 0{
		log.Fatal("Evironment variable DATABASE_URL is empty")
	}
}

func (db database) Connect()(*sql.DB, error){
	conn, err := sql.Open("postgres",url)
	return conn, err
}

type dbyer interface{
	QueryAll(Conn *sq.DB)
}
