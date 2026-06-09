package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type TimeResponse struct {
	Time     string `json:"time"`
	UnixTime int64  `json:"unix_time"`
	TimeZone string `json:"timezone"`
}

func getCurrentTime(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	zone, _ := now.Zone()

	response := TimeResponse{
		Time:     now.Format(time.RFC3339),
		UnixTime: now.Unix(),
		TimeZone: zone,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/time", getCurrentTime)

	fmt.Println("Server running on http://localhost:8080")
	fmt.Println("GET http://localhost:8080/time")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
