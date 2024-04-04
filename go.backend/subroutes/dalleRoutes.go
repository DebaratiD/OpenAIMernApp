package subroutes

import "github.com/gin-gonic/gin"

func CreateImage(c *gin.Context) {
	c.String(200, "Hello from Stable Difussion!")
}
