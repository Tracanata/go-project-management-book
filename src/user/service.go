package user

import (
	"fmt"
	"math/rand/v2"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) RegisterUser(user *User) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPass)
	user.Username = fmt.Sprintf("US%03d", rand.IntN(999)+1)
	err = s.repo.Register(user)
	if err != nil {
		return err
	}
	return nil
}
