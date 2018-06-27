package history

import (
	"fmt"
	"strconv"
	"github.com/dyninc/qstring"
)

type TypeQuery struct {
	ID   []string
	View string
}

func (h *TypeApprovedHistory) NewWorker() {
	for current := range h.ChanDate {

		fmt.Println(current)

		page := 0

		for {
			page++

			res, errBase := h.BaseRequest(strconv.Itoa(page), current)
			if errBase != nil {
				fmt.Println(errBase)
			}

			if len(res.Data) == 0 {
				break
			}

			listID := makeListID(res)
			query := makeQuery(listID)

			full, errFull := h.FullRequest(query)
			if errFull != nil {
				fmt.Println(errFull)
			}

			fmt.Println(full)
		}

		h.SyncGroup.Done()
	}
}

func (h *TypeApprovedHistory) CreateWorker() {
	for i := 0; i < 10; i++ {
		go h.NewWorker()
	}
}

func makeListID(res *BaseRespType) []string {
	var listID []string
	{
	}
	for _, image := range res.Data {
		listID = append(listID, image.Media_id)
	}
	return listID
}

func makeQuery(listID []string) string {
	query := &TypeQuery{
		ID:   listID,
		View: "full",
	}
	q, errQ := qstring.MarshalString(query)
	if errQ != nil {
		fmt.Println(errQ)
	}

	return q
}
