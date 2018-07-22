package main

import (
	"github.com/gin-gonic/gin"
	"github.com/my-stocks-pro/approved-history-service/history"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

func main() {
	router := gin.Default()

	h := history.New()

	fmt.Println(h.Config)

	router.GET("history/approved", func(c *gin.Context) {

		var dateRange history.TypeDateRange

		body := c.Request.Body
		b, e := ioutil.ReadAll(body)
		if e != nil {
			fmt.Println(e)
		}

		err := json.Unmarshal(b, &dateRange)
		if err != nil {
			fmt.Println(err)
		}

		//h := history.New(dateRange)

		h.Start = h.GetDate(dateRange.Start)
		h.End = h.GetDate(dateRange.End)

		go h.CreateWorker()

		h.CreateTasks()

		h.SyncGroup.Wait()
		h.SyncGroupPost.Wait()

		c.String(200, "Success")
	})

	router.Run(":8002")
}
