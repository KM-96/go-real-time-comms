package poll

import (
	"encoding/json"
	"fmt"
	"go-comms/common/model"
	"go-comms/common/util"
	"net/http"
	"net/url"
	"time"
)

const (
	PollDurationInSeconds = 5
)

func Poll() {
	ticker := time.NewTicker(PollDurationInSeconds * time.Second)
	defer ticker.Stop()

	for {
		// Wait for the ticker to signal that it's time to poll
		<-ticker.C

		// Perform the HTTP GET request to the /messages endpoint
		afterTime := util.GetCurrentTime().Add(-5 * time.Second).Format(time.DateTime)
		queryParams := url.Values{}
		queryParams.Add("after", afterTime)

		url := "http://localhost:8080/messages" + "?" + queryParams.Encode()
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error making HTTP request:", err)
			continue
		}
		defer resp.Body.Close()

		// Check the HTTP status code
		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Received non-OK status code: %d\n", resp.StatusCode)
			continue
		}

		// Decode the JSON response into a slice of Message structs
		var messages []model.Message
		if err := json.NewDecoder(resp.Body).Decode(&messages); err != nil {
			fmt.Println("Error decoding JSON response:", err)
			continue
		}

		// Print or process the received messages
		fmt.Println("Received Messages:", messages)
	}
}
