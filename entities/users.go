package entities

import "time"

type User struct {
	Id        int       `json:"id" db:"id"`
	FirstName string    `json:"firstname" db:"first_name"`
	LastName  string    `json:"lastname" db:"last_name"`
	EmailId   string    `json:"emailid" db:"email_id"`
	Password  string    `json:"password" db:"password"`
	CreatedAt time.Time `json:"createdat" db:"created_at"`
	UpdatedAt time.Time `json:"updatedat" db:"updated_at"`
}
