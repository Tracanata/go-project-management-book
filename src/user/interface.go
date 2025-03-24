package user

type UserRepository interface {
	Register(user *User) error
	Login(username string) (*User, error)
}

type UserService interface {
	RegisterUser(user *User) error
	LoginUser(user, password string) (*RespSuccessLogin, error)
}
