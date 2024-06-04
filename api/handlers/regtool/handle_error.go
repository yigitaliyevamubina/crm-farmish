package v1

import (
	"crm-farmish/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func HandleError(c *gin.Context, err error, l *zap.Logger, statusCode int, msg string) bool {
	if err == nil {
		return false
	}
	c.JSON(statusCode,
		&models.ResponseError{
			Code:    http.StatusText(statusCode),
			Data:    msg,
			Message: err.Error(),
		})
	l.Log(1, err.Error())
	return true
}
