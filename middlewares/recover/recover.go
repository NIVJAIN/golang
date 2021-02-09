package recover

import (
	"fmt"
	// "ginDemo/common/alarm"
	"github.com/gin-gonic/gin"
	"github.com/nivjain/7-ginInterfaceMongoDBRabbitMQ-JWT-RTL/alarm"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				alarm.Panic(fmt.Sprintf("%s", r))
			}
		}()
		c.Next()
	}
}
