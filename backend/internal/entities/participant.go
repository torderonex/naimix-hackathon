package entities

type Participant struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	TeamID     int    `json:"team_id" db:"team_id"`
	Role       string `json:"role"`
	BirthDate  string `json:"birthdate"`
	BirthPlace string `json:"birthplace"`
}
