package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetApi(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "GetApi",
	})
}

func PostApi(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "PostApi",
	})
}
