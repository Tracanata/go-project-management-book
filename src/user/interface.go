package user

type UserRepository interface {
	RegisterUser(user User) error
}

type UserService interface {
	Register(user User) error
}
