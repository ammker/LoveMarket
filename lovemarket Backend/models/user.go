package models

type User struct {
	UserID   int64  `gorm:"column:userid" db:"userId"`
	Username string `gorm:"column:username" db:"userName"`
	Password string `gorm:"column:password" db:"password"`
	Email    string `db:"email"`
	Token    string
}
