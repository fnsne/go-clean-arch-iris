package main

import (
	_bookHttpDelivery "clean-architecture-template/services/books/delivery/http"
	_bookRepo "clean-architecture-template/services/books/repository/postgres"
	_bookUsecase "clean-architecture-template/services/books/usecase"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/kataras/iris/v12"
	"log"
)

//learning from https://github.com/bxcodec/go-clean-arch/blob/master/app/main.go

func main() {
	//initial db connection
	pgConn, err := pgx.Connect(context.Background(), "postgres://fnsne:@localhost:5432/mydb")
	if err != nil {
		log.Fatal("failed to connect to database, error=", err)
	}
	defer pgConn.Close(context.Background())

	//initial iris app
	app := iris.Default()

	//直接用sqlc產生的queries來當作repository，等到真的要寫測試的時候，再來抽出repo的interface也不遲。
	booksRepo := _bookRepo.New(pgConn)
	bookUsecase := _bookUsecase.NewBookUsecase(booksRepo)
	_bookHttpDelivery.AddBookHandlersToApp(app, bookUsecase)

	//start server
	err = app.Listen(":8080")
	if err != nil {
		fmt.Println("running iris server failed, error=", err)
		return
	}
}
