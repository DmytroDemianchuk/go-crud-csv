package v1

import (
	"github.com/dmytrodemianchuk/go-crud-csv/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func newErrorResponse(c *gin.Context, statusCode int, err error) {
	logrus.Error(err)
	c.AbortWithStatusJSON(statusCode, domain.ErrorResponse{Error: err.Error()})
}
