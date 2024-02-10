package model

import "time"

type GetElectionPeriodResponse struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}
