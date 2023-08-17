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

	//把repo, usecase,delivery都在main中初始化，這樣就相依性就會很明確，也很好測試。
	//直接用sqlc產生的queries來當作repository，等到真的要寫測試的時候，再來抽出repo的interface也不遲。
	booksRepo := _bookRepo.New(pgConn)
	bookUsecase := _bookUsecase.NewBookUsecase(booksRepo)
	//每個delivery把handler掛在app的行為分開來，這樣未來要測試的時候就可以不用接上所有的delivery，只用接需要測試的，就可以測試整個httpRequest了。
	_bookHttpDelivery.AddBookHandlersToApp(app, bookUsecase)

	//start server
	err = app.Listen(":8080")
	if err != nil {
		fmt.Println("running iris server failed, error=", err)
		return
	}
}
