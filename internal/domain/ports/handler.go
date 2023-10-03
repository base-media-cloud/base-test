package ports

import "github.com/gin-gonic/gin"

type Handler interface {
	Get(c *gin.Context)
}
