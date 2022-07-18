package tblBkng

type Cafe struct {
	ID       int    `json:"id"`
	CafeName string `json:"cafe_name"`
	AvgTime  int    `json:"avg_time"`
	AvgCheck int    `json:"avg_check"`
}
