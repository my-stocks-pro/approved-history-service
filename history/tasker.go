package history

import (
	"time"
	"strconv"
	"fmt"
)

func (h *TypeApprovedHistory) NewDate(t time.Time, timestamp int64) *TypeCurrDate {
	return &TypeCurrDate{
		TimeStamp: t,
		UnixDate:  timestamp,
		DateStr:   t.Format("2006-01-02"),
		Day:       h.CheckDate(strconv.Itoa(time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, time.UTC).Day())),
		Month:     h.CheckDate(fmt.Sprintf("%d", t.Month())),
		Year:      strconv.Itoa(t.Year()),
	}
}

func (h *TypeApprovedHistory) CreateTasks() {

	newMonth := 0
	for timestamp := h.Start.UnixDate; timestamp <= h.End.UnixDate; timestamp += h.OneDay {
		t := time.Unix(timestamp, 0)
		tmpMonth := int(t.Month())
		if tmpMonth != newMonth {

			h.SyncGroup.Add(1)

			newMonth = tmpMonth

			h.ChanDate <- h.NewDate(t, timestamp)
		}
	}
}
