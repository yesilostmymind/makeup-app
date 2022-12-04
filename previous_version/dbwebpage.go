package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//when using local host please includ sslmode=disable in arguments for .Open, otherwise exclude

func main() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=makeupdbtest password=wrapitup45 sslmode=disable")

//when i frist wrote this, there was a pw typo, and panic triggered, so i'm not sure why this isn't working 

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

/* creates a database called "database" and then pings it haven't been able to see it in pgAdmin
	
	database := db.DB()

	err = database.Ping()
*/

//say cosmetics is undefined ??

	db.First(&cosmetics)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connection to PostgreSQL sucess!");

}