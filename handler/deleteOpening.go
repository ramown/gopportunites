package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{
		"msg": "DELETE Opening",
	})
}
