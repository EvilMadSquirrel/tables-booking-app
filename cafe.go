package tblBkng

type Cafe struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	AvgTime  int    `json:"avg_time"`
	AvgCheck int    `json:"avg_check"`
}
