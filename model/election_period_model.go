package model

import "time"

type SetElectionPeriodRequest struct {
	Start string `json:"start" binding:"required"`
	End   string `json:"end" binding:"required"`
}

type GetElectionPeriodResponse struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}
