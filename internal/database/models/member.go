package models

// Member represents a gym member
type Member struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Age        int    `json:"age"`
	JoinedDate string `json:"joined_date"`
}
