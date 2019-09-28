package echo

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//HandleGet return
func HandleGet(log *logrus.Logger) gin.HandlerFunc {

	return func(c *gin.Context) {
		message := c.Query("message")
		log.Info(message)
		c.JSON(200, gin.H{
			"message": message,
		})
	}

}
