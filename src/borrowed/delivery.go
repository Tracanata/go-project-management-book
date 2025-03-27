package borrowed

import (
	"go-project-management-book/helper"

	"github.com/gin-gonic/gin"
)

type BorrowedHandler struct {
	service BorrowedService
}

func NewBorrowedHandler(service BorrowedService) *BorrowedHandler {
	return &BorrowedHandler{service: service}
}

func (dlv *BorrowedHandler) SetBorrowBook(ctx *gin.Context) {
	var req ReqBorrowedBook
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.SendResponseError(ctx, helper.ErrBadRequest)
		return
	}

	err := dlv.service.BorrowBook(req.UserId, req.BookId)
	if err != nil {

	}
}
