package models

type User struct {
	ID      int64   `json:"id"`
	Fname   string  `json:"fname"`
	City    string  `json:"city"`
	Phone   int64   `json:"phone"`
	Height  float32 `json:"height"`
	Married bool    `json:"married"`
}
