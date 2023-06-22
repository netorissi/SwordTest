package model

type Session struct {
	ID          int    `json:"id" db:"id"`
	UserID      int    `json:"user_id" db:"user_id"`
	AccessToken string `json:"access_token" db:"access_token"`
}
