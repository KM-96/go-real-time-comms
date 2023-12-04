package util

import (
	"fmt"
	"time"
)

func GetCurrentTime() time.Time {
	current := time.Now()
	formattedTime := current.Format(time.DateTime)
	formattedCurrentTime, err := time.Parse(time.DateTime, formattedTime)
	if err != nil {
		fmt.Println(err)
		return time.Now() // Need to change this
	}
	return formattedCurrentTime
}
