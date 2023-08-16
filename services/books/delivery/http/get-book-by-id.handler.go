package http

import (
	"github.com/kataras/iris/v12"
	"strconv"
)

func (h *BookHandlers) GetBookByID(ctx iris.Context) {
	idStr := ctx.Params().Get("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_ = ctx.JSON("cannot parse input book id in query")
		return
	}
	book, err := h.bookUseCase.GetBookByID(id)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		_ = ctx.JSON("error when getting book by id")
		return
	}
	_ = ctx.JSON(book)
}
