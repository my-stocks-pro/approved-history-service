package history

import (
	"fmt"
)

func (h *TypeApprovedHistory) NewWorker() {
	for current := range h.ChanDate {

		fmt.Println(current)

		h.SyncGroup.Done()
	}
}


func (h *TypeApprovedHistory) CreateWorker() {
	for i := 0; i < 10; i++ {
		go h.NewWorker()
	}
}

