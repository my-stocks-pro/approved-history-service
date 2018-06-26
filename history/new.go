package history

import (
	"time"
	"strconv"
	"fmt"
)

type TypeStartDate struct {
	TimeStamp time.Time
	Date      string
	Year      string
	Month     string
	Day       string
}

type TypeCurrentDate struct {
	TimeStamp time.Time
	Date      string
	Year      string
	Month     string
	Day       string
}

type TypeApprovedHistory struct {
	Start   TypeStartDate
	Current TypeCurrentDate
}

func New() *TypeApprovedHistory {
	return &TypeApprovedHistory{
		//os.Getenv("START"),
		TypeStartDate{},
		TypeCurrentDate{},
	}
}

func (h *TypeApprovedHistory) GetCurrDate() {
	currDate := time.Now()

	h.Current.TimeStamp = currDate
	h.Current.Date = currDate.Format("2006-01-02")
	h.Current.Year = strconv.Itoa(currDate.Year())
	h.Current.Month = CheckDate(fmt.Sprintf("%d", currDate.Month()))
	h.Current.Day = CheckDate(strconv.Itoa(currDate.Day()))
}

func CheckDate(tmp string) string {
	var res string
	if len(tmp) == 1 {
		res = "0" + tmp
	} else {
		res = tmp
	}
	return res
}
