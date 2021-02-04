package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/vibhugarg123/Go_1/config"
	"github.com/vibhugarg123/Go_1/entities"
	"time"
)

// mysql connection, so that during bootstrap it gives err, rest it returns static object
type UserService interface {
	Add(user entities.User) (entities.User, error)
}

type userService struct {
	Db *sqlx.DB
}

func NewUserService() UserService {
	db, _ := config.GetMySqlConnection()
	return &userService{Db: db}
}

func (u userService) Add(user entities.User) (entities.User, error) {
	tx := u.Db.MustBegin()
	u.Db.MustExec("INSERT INTO users  (first_name,last_name,email_id,password,created_at,updated_at) VALUES (?,?,?,?,?,?)", user.FirstName, user.LastName, user.EmailId, user.Password, time.Now(), time.Now())
	err := tx.Commit()
	return user, err
}
