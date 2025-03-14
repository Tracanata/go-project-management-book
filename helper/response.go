package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendResponseSuccess(c *gin.Context, resp Response) {
	c.JSON(http.StatusOK, resp)
}

func SendResponseError(c *gin.Context, err ErrorStruct) {
	resp := Response{
		ResponseCode: err.Code,
		Description:  err.Description,
		Message:      err.Message,
	}
	c.JSON(err.HTTPCode, resp)
}
