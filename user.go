package tblBkng

type User struct {
	ID       int    `json:"id"`
	UserName string `json:"user_name"`
	Phone    int    `json:"phone"`
	Password string `json:"password"`
}
