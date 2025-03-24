package user

import "time"

type User struct {
	Id         int       `json:"id"`
	Username   string    `json:"username"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Role       string    `json:"role" binding:"required,oneof=admin member"`
	Created_at time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func (User) TableName() string {
	return "user"
}
