package model

type Link struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Url        string `json:"url"`
	UserID     int    `json:"user_id"`
	Background string `json:"background"`
	ClickCount int    `json:"click_count"`
}
