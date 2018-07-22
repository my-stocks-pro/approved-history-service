package history

import (
	"time"
	"fmt"
	"sync"
)

type TypeDateRange struct {
	Start string `form:"start"`
	End   string `form:"end"`
}

type TypeCurrDate struct {
	TimeStamp time.Time
	UnixDate  int64
	DateStr   string
	Day       string
	Month     string
	Year      string
}

type TypeDate struct {
	TimeStamp time.Time
	UnixDate  int64
	DateStr   string
}

type TypeApprovedHistory struct {
	Config        *TypeConfig
	Start         *TypeDate
	End           *TypeDate
	ChanDate      chan *TypeCurrDate
	ChanPost      chan *DataImageType
	OneDay        int64
	SyncGroup     sync.WaitGroup
	SyncGroupPost sync.WaitGroup
}

func New() *TypeApprovedHistory {
	h := &TypeApprovedHistory{
		Config:   LoadConfig(),
		Start:    &TypeDate{},
		End:      &TypeDate{},
		ChanDate: make(chan *TypeCurrDate),
		ChanPost: make(chan *DataImageType),
		OneDay:   int64(86400), // a day in seconds.
	}

	return h
}

func (h *TypeApprovedHistory) GetDate(timeStr string) *TypeDate {
	var t time.Time
	var err error

	if len(timeStr) != 0 {
		t, err = time.Parse("2006-01-02", timeStr)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		t = time.Now()
	}

	return &TypeDate{
		t,
		t.Unix(),
		timeStr,
	}
}

func (h *TypeApprovedHistory) CheckDate(tmp string) string {
	var res string
	if len(tmp) == 1 {
		res = "0" + tmp
	} else {
		res = tmp
	}
	return res
}
