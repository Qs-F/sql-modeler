//go:generate sql-modeler -o mods -sql schema.sql -pkg mods
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/Qs-F/sql-modeler/_example/mods"
	_ "github.com/mattn/go-sqlite3"
	"github.com/volatiletech/sqlboiler/boil"
)

func main() {
	db, err := sql.Open("sqlite3", "models.sqlite")
	if err != nil {
		log.Fatal(err)
	}

	// select Users
	users, err := mods.Users().All(context.Background(), db)
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range users {
		fmt.Println(user.Name)
	}

	// begin transaction
	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	// insert new User
	user := &mods.User{}
	user.Name = "mattn"
	err = user.Insert(context.Background(), tx, boil.Infer())
	if err != nil {
		log.Fatal(err)
	}

	tx.Commit()

	user, err = mods.Users().One(context.Background(), db)
	if err != nil {
		log.Fatal(err)
	}
}
