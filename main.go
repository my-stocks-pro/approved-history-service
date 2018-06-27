package main

import (
	"github.com/my-stocks-pro/approved-history-service/history"
)


func main() {
	h := history.New()

	go h.CreateWorker()

	h.CreateTasks()

	h.SyncGroup.Wait()
	h.SyncGroupPost.Wait()
}
