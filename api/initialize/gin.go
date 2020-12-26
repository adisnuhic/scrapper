package initialize

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

// Gin returns instance of gin
func Gin() *gin.Engine {
	g := gin.Default()
	g.Use(requestid.New()) // use request id
	return g
}
