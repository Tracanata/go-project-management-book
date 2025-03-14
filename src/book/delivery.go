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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.AddBook(book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
