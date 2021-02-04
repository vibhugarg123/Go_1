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

var UsersSchema = `
CREATE TABLE users ( 
		id integer,
		first_name text,
  		last_name text,
    	email_id text,
    	password text,
   	    created_at DATETIME,
        updated_at DATETIME
   );
`
