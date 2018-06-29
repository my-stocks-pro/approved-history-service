package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

type Date struct {
	Start string `form:"start"`
	End   string `form:"end"`
}


func bindDate(c *gin.Context) {
	var dateRange Date
	if c.Bind(&dateRange) == nil {
		fmt.Println(dateRange.Start)
		fmt.Println(dateRange.End)
	}
	c.String(200, "Success")
}


func main() {
	router := gin.Default()

	router.GET("history/approved", bindDate)


	//{
	//
	//	if c.BindJSON(data) == nil {
	//		fmt.Println("Post fron scheduler -> ", data)
	//	}
	//})

	//h := history.New()
	//
	//go h.CreateWorker()
	//
	//h.CreateTasks()
	//
	//h.SyncGroup.Wait()
	//h.SyncGroupPost.Wait()

	router.Run(":8002")
}
