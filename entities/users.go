package entities

import "time"

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

type User struct {
	Id        int       `db:"id"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	EmailId   string    `db:"email_id"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
