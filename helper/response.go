package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendResponseSuccess(c *gin.Context, resp Response) {
	c.JSON(http.StatusOK, resp)
}
