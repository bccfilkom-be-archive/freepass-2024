package domain

import "time"

type ElectionTimes struct {
	ID        int       `json:"id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

type ElectionTimesRequest struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}
