package user

import (
	"errors"
	"gorm.io/driver/postgres"
	"golang.org/x/crypto/bcrypt"
)

//User ...
type OmsUser struct {
	ID        int64  `db:"id, primarykey, autoincrement" json:"id"`
	Email     string `db:"email" json:"email"`
	Password  string `db:"password" json:"-"`
	UserName  string `db:"username" json:"username"`
}
