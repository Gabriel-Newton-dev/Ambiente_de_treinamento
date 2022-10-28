package main

import (
	"log"

	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open("mysql", "golang:G1ogo@2060@localhost/devbook")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
