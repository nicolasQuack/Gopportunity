package router

import (
	g "github.com/gin-gonic/gin"
	h "github.com/nicolasQuack/Gopportunity/handlers"
)

// InitRoutes initializes the routes
func initRoutes(r *g.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/opening", h.GetOpeningHandler)
		v1.POST("/opening", h.CreateOpeningHandler)
		v1.DELETE("/opening", h.DeleteOpeningHandler)
		v1.PUT("/opening", h.UpdateOpeningHandler)
		v1.GET("/openings", h.GetOpeningsHandler)
	}
}