package http

import (
	_bookUsecase "clean-architecture-template/services/books/usecase"
	"github.com/kataras/iris/v12"
)

type BookHandlers struct {
	bookUseCase *_bookUsecase.BookUsecase
}

func AddBookHandlersToApp(app *iris.Application, bookUseCase *_bookUsecase.BookUsecase) {
	handlers := BookHandlers{bookUseCase: bookUseCase}
	api := app.Party("api/books")
	api.Get("/{id:uint64}", handlers.GetBookByID)
}
