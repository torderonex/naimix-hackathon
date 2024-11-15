package entities

type Team struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	UserID int    `json:"user_id"`
}
