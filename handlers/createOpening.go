package handlers

import (
	"net/http"

	g "github.com/gin-gonic/gin"
)

// CreateOpeningHandler handles the POST /api/v1/opening route
func CreateOpeningHandler(ctx *g.Context) {
	ctx.JSON(http.StatusOK, g.H{
		"message": "Post Opening soon!",
	})
}