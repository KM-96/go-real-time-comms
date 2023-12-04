package handlers

import (
	"encoding/json"
	"fmt"
	"go-comms/service/pkg/store"
	"net/http"
	"time"
)

const (
	QueryParamAfter = "after"
)

func GetMessagesHandler(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header to JSON
	w.Header().Set("Content-Type", "application/json")

	// Encode the messages slice to JSON and write it to the response
	var err error
	after := r.URL.Query().Get(QueryParamAfter)
	if after == "" {
		err = json.NewEncoder(w).Encode(store.GetAllMessages())
	} else {
		t, parseErr := time.Parse(time.DateTime, after)
		if parseErr == nil {
			err = json.NewEncoder(w).Encode(store.GetMessagesAfter(t))
		} else {
			fmt.Println(parseErr)
			http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
			return
		}
	}
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}
}
