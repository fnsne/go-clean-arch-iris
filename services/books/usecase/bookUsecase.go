package usecase

import (
	books "clean-architecture-template/services/books/repository/postgres"
	"context"
)

type BookUsecase struct {
	bookRepo *books.Queries
}

func (u *BookUsecase) GetBookByID(id uint64) (books.Book, error) {
	return u.bookRepo.GetBook(context.Background(), int32(id))
}

func NewBookUsecase(bookRepo *books.Queries) *BookUsecase {
	return &BookUsecase{bookRepo: bookRepo}
}
