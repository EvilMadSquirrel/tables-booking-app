package tblBkng

import "time"

type Booking struct {
	ID      int       `json:"id"`
	Start   time.Time `json:"start"`
	End     time.Time `json:"end"`
	TableID int       `json:"table_id"`
}
