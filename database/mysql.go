package database

import (
	// database driver
	_ "github.com/go-sql-driver/mysql"
	sql "github.com/jmoiron/sqlx"
	"log"
)

type Config struct {
	Addr 		string
	DBName   	string
	UserName 	string
	Password 	string
	Config		string
}

func (c *Config) Dsn() string {
	return c.UserName + ":" + c.Password + "@tcp(" + c.Addr + ")/" + c.DBName + "?" + c.Config
}

func NewMySQL(c *Config) (db *sql.DB){
	db, err := sql.Connect("mysql", c.Dsn())
	if err != nil {
		log.Fatalf("open mysql error(%v)", err)
		panic(err)
	}

	return
}
