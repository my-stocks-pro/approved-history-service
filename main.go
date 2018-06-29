package main

import (
	"github.com/gin-gonic/gin"
	"github.com/my-stocks-pro/approved-history-service/history"
	"fmt"
)

func main() {
	router := gin.Default()

	router.GET("history/approved", func(c *gin.Context) {

		var dateRange history.TpeDateRange
		err := c.Bind(&dateRange)
		if err != nil {
			fmt.Println(err)
		}

		h := history.New(dateRange)

		go h.CreateWorker()

		h.CreateTasks()

		h.SyncGroup.Wait()
		h.SyncGroupPost.Wait()

		c.String(200, "Success")
	})

	router.Run(":8002")
}
