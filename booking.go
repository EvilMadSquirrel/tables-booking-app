package tblBkng

import "time"

type Booking struct {
	ID        int       `json:"id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	TableID   int       `json:"table_id"`
	UserID    int       `json:"user_id"`
}
