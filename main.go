package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	t1 := time.Now()
	t2 := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC) // 2026 year, 1 month, 1 day, 0 hour, 0 minute, 0 seconds, 0 nanoseconds, time.UTC
	t3 := t2.Sub(t1) / 86400000000000

	r := gin.Default()
	r.GET("/date", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"days": t3,
		})
	})
	fmt.Println(t2.Sub(t1))
	r.Run(":8000")
}
