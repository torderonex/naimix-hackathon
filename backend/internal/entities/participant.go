package entities

import "time"

type Participant struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	TeamName   string `json:"team_name" db:"team_name"`
	TeamID     int       `json:"team_id" db:"team_id"`
	Role       string    `json:"role"`
	BirthDate  time.Time `json:"birthdate"`
	BirthPlace string    `json:"birthplace"`
}
