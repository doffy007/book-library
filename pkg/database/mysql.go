package database

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type ConfigDB struct {
	Host     string
	User     string
	Password string
	Name     string
}

func Connect(cnf ConfigDB) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		cnf.User,
		cnf.Password,
		cnf.Host,
		cnf.Name,
	)
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connection")
	}

	return db, err
}
