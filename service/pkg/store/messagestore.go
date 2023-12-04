package store

import (
	"fmt"
	"go-comms/common/model"
	"go-comms/common/util"
	"time"

	"github.com/google/uuid"
)

var messages []model.Message

func GetAllMessages() *[]model.Message {
	return &messages
}

func GetMessagesAfter(t time.Time) *[]model.Message {
	var result []model.Message
	for _, msg := range messages {
		if msg.Timestamp.After(t) {
			result = append(result, msg)
		}
	}
	return &result
}

func AddMessage(m *model.Message) {
	messages = append(messages, *m)
}

func AddInput(input string) {
	if input != "" {
		message := model.Message{
			Id:        uuid.NewString(),
			Content:   input,
			Timestamp: util.GetCurrentTime(),
		}
		AddMessage(&message)
		PrintMessages()
	}
}

func PrintMessages() {
	for _, m := range messages {
		fmt.Printf("%+v", m)
		fmt.Println()
	}
}
