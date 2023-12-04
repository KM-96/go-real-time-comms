package model

import "time"

type Message struct {
	Id        string
	Content   string
	Timestamp time.Time
}
