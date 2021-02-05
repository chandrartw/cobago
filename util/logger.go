package util

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] %s %s %d %s \n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC822),
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
		)
	})
}

func SetupLoggerOutput() {
	fmt.Println("oke")
	f, _ := os.Create(os.Getenv("LOG_DIR"))
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}