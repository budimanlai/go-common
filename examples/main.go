package main

import (
	"fmt"

	common "github.com/budimanlai/go-common"
	models "github.com/budimanlai/go-common/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func main() {
	// Your code here
	hostname := "localhost"
	username := "root"
	password := ""
	database := "emove"
	port := "3306"

	var err error
	common.Db, err = sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", username, password, hostname, port, database))
	if err != nil {
		panic(err)
	}

	var user *models.User
	user, err = models.FindUserByUsername(common.Db, "superadmin")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("User found:", user)
	}
	common.Db.Close()

}
