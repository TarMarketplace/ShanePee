package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"shanepee.com/api/apperror"
)

func handleError(c *gin.Context, err apperror.AppError) {
  if err.ShouldLog() {
    // TODO: use proper logger
    fmt.Println("Internal server error: ", err.Cause())
  }
  c.JSON(err.Code(), gin.H{
    "message": err.Message(),
  })
}
