package handler

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(g *gin.RouterGroup) {
	g.GET("/hello", hello)
}

// @BasePath /api/v1

// hello godoc
// @Summary hello example
// @Schemes
// @Description hello
// @Tags hello
// @Accept json
// @Produce json
// @Success 200 {string} ok
// @Router /hello [get]
func hello(c *gin.Context) {
	c.JSON(200, "ok")
}
