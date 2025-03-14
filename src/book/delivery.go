package book

import (
	"go-project-management-book/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

func (h *BookHandler) GetBookById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := h.service.GetBookById(id)
	if err != nil {
		helper.SendResponseError(c, helper.ErrNotFound)
		c.JSON(http.StatusNotFound, gin.H{"message": "Book Not Found"})
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
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		helper.SendResponseError(c, helper.ErrBadRequest)
		return
	}
	if err := h.service.UpdateBook(book); err != nil {
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
	id, err := strconv.Atoi(c.Param("id"))
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

	err = h.service.DeleteBook(id)
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
