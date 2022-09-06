package route

import (
	"github.com/gin-gonic/gin"
	"homepage-server/logic"
	"net/http"
)

func Route(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/fetchNbaSeason", func(c *gin.Context) {
		c.JSON(http.StatusOK, logic.FetchNbaSeason(c))
	})
}
