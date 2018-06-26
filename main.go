package main

import (
	"github.com/my-stocks-pro/approved-history-service/history"
	"fmt"
	"time"
)

func main() {
	historyClient := history.New()


	start := "2016-08-01"
	t, err := time.Parse("2006-01-02", start)
	if err != nil {
		fmt.Println(err)
	}

	startDate := t.Unix()

	oneDay := int64(86400) // a day in seconds.

	endDate := time.Now().Unix()

	for timestamp := startDate; timestamp <= endDate; timestamp += oneDay {
		t := time.Unix(timestamp, 0)
		fmt.Println(t)
	}


	fmt.Println(historyClient)
}
