package user

type UserRepository interface {
	Register(user *User) error
}

type UserService interface {
	RegisterUser(user *User) error
}
