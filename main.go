package main

import (
	"github.com/gin-gonic/gin"
	"github.com/my-stocks-pro/approved-history-service/history"
	"fmt"
)

type Date struct {
	Start string `form:"start"`
	End   string `form:"end"`
}

func main() {
	router := gin.Default()

	router.GET("history/approved", func(c *gin.Context) {

		h := history.New()

		go h.CreateWorker()

		h.CreateTasks()

		var dateRange Date

		err := c.Bind(&dateRange)
		if err != nil {
			fmt.Println(err)
		}

		h.SyncGroup.Wait()
		h.SyncGroupPost.Wait()

		c.String(200, "Success")
	})

	router.Run(":8002")
}
