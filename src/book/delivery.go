package book

import (
	"errors"
	"go-project-management-book/helper"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookHandler struct {
	service BookService
}

func NewBookHandler(service BookService) *BookHandler {
	return &BookHandler{service: service}
}

func (h *BookHandler) GetAllBooks(c *gin.Context) {
	books, err := h.service.GetAllBooks()
	if err != nil {
		resp := helper.ErrorStruct{
			Code:        helper.RCGeneralError,
			HTTPCode:    http.StatusInternalServerError,
			Description: helper.DescriptionFailed,
			Message:     err.Error(),
		}
		helper.SendResponseError(c, resp)
		return
	}

	resp := helper.Response{
		ResponseCode: helper.RCSuccess,
		Description:  helper.DescriptionSuccess,
		Message:      http.StatusText(http.StatusOK),
		Data:         books,
	}
	helper.SendResponseSuccess(c, resp)
}

func (h *BookHandler) GetBookByCodeBook(c *gin.Context) {
	code := c.Param("codeBook")
	book, err := h.service.GetBookByCodeBook(code)
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.SendResponseError(c, helper.ErrNotFound)
			return
		}

		resp := helper.ErrorStruct{
			Code:        helper.RCGeneralError,
			HTTPCode:    http.StatusInternalServerError,
			Description: helper.DescErrorGeneral,
			Message:     err.Error(),
		}
		helper.SendResponseError(c, resp)

	}
	resp := helper.Response{
		ResponseCode: helper.RCSuccess,
		Description:  helper.DescriptionSuccess,
		Message:      http.StatusText(http.StatusOK),
		Data:         book,
	}
	helper.SendResponseSuccess(c, resp)
}

func (h *BookHandler) AddBook(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		helper.SendResponseError(c, helper.ErrBadRequest)
		return
	}
	if err := h.service.AddBook(book); err != nil {
		resp := helper.ErrorStruct{
			Code:        helper.RCGeneralError,
			HTTPCode:    http.StatusInternalServerError,
			Description: helper.DescErrorGeneral,
			Message:     err.Error(),
		}
		helper.SendResponseError(c, resp)
		return
	}
	resp := helper.Response{
		ResponseCode: helper.RCSuccess,
		Description:  helper.DescriptionSuccess,
		Message:      http.StatusText(http.StatusOK),
		Data:         nil,
	}

	helper.SendResponseSuccess(c, resp)
}

func (h *BookHandler) UpdateBook(c *gin.Context) {
	code := c.Param("codeBook")

	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		helper.SendResponseError(c, helper.ErrBadRequest)
		return
	}

	book.Code_Book = code
	err := h.service.UpdateBook(book)
	if err != nil {
		resp := helper.ErrorStruct{
			Code:        helper.RCGeneralError,
			HTTPCode:    http.StatusInternalServerError,
			Description: helper.DescErrorGeneral,
			Message:     err.Error(),
		}
		helper.SendResponseError(c, resp)
		return
	}
	resp := helper.Response{
		ResponseCode: helper.RCSuccess,
		Description:  helper.DescriptionSuccess,
		Message:      http.StatusText(http.StatusOK),
		Data:         nil,
	}

	helper.SendResponseSuccess(c, resp)
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
	code := c.Param("codeBook")

	err := h.service.DeleteBook(code)
	if err != nil {
		helper.SendResponseError(c, helper.ErrNotFound)
		return
	}

	resp := helper.Response{
		ResponseCode: helper.RCSuccess,
		Description:  helper.DescriptionSuccess,
		Message:      http.StatusText(http.StatusOK),
		Data:         nil,
	}
	helper.SendResponseSuccess(c, resp)
}
