package router

import (
	"net/http"

	g "github.com/gin-gonic/gin"
)

// InitRoutes initializes the routes
func initRoutes(r *g.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/opening", func(ctx *g.Context) {
			ctx.JSON(http.StatusOK, g.H{
				"message": "Get Opening soon!",
			})
		})
		v1.POST("/opening", func(ctx *g.Context) {
			ctx.JSON(http.StatusOK, g.H{
				"message": "Post Opening soon!",
			})
		})
		v1.DELETE("/opening", func(ctx *g.Context) {
			ctx.JSON(http.StatusOK, g.H{
				"message": "Delete Opening soon!",
			})
		})
		v1.PUT("/opening", func(ctx *g.Context) {
			ctx.JSON(http.StatusOK, g.H{
				"message": "Put Opening soon!",
			})
		})
		v1.GET("/openings", func(ctx *g.Context) {
			ctx.JSON(http.StatusOK, g.H{
				"message": "Get Openings soon!",
			})
		})
		
	}
}