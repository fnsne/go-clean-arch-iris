package main

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattes/migrate/source/file"
)

func main() {
	m, err := migrate.New(
		"file://migrations",
		"pgx5://fnsne:@localhost:5432/mydb")
	if err != nil {
		fmt.Println("err=", err)
		panic(err)
	}
	err = m.Up()
	fmt.Println("migration err=", err)

}
