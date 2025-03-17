package user

type User struct {
	Id         int    `json:"id"`
	Name       int    `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Role       string `json:"role"`
	Created_at string `json:"createdAt"`
}

func (User) TableName() string {
	return "user"
}
