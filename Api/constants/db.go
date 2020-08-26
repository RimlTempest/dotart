package constants

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// Init database.
func Init() *gorm.DB {
	DBMS := "mysql"
	USER := "team_a"
	PASS := "testtest"
	PROTOCOL := "tcp(mysql:3306)" //← ここのmysqlはサービス名です
	DBNAME := "dotart"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		fmt.Println("CONNECTION : ERROR")
		panic(err.Error())
	}
	fmt.Println("CONNECTION : OK")
	db.LogMode(true)

	DB = db
	return DB
}

// GetDB returns the connection to the database.
func GetDB() *gorm.DB {
	return DB
}
