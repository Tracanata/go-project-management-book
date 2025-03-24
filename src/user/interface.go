package user

type UserRepository interface {
	Register(user *User) error
	Login(username string) (*User, error)
	GetDataUser(username string) (*User, error)
}

type UserService interface {
	RegisterUser(user *User) error
	LoginUser(user, password string) (*RespSuccessLogin, error)
	GetProfile(username string) (*RespSuccessLogin, error)
}
