package middlieware

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func RequestShow() gin.HandlerFunc {
	logFilePath := "request.log" // 设置日志文件路径
	logFile, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	logger := log.New(logFile, "", log.LstdFlags)

	return func(c *gin.Context) {

		c.Next()

		clientIP := c.ClientIP()
		request := c.Request

		text := fmt.Sprintf("From: %v\nHeader:\n", clientIP)
		for i, v := range request.Header {
			text = text + fmt.Sprintf("%v : %v\n", i, v)
		}

		text = text + fmt.Sprintf("Body:\n %v\n", request.Body)

		logger.Println(text)

	}
}
