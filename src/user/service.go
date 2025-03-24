package user

import (
	"errors"
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

func (svc *userService) RegisterUser(user *User) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPass)
	user.Username = fmt.Sprintf("US%03d", rand.IntN(999)+1)
	err = svc.repo.Register(user)
	if err != nil {
		return err
	}
	return nil
}

func (svc *userService) LoginUser(username, password string) (*RespSuccessLogin, error) {
	user, err := svc.repo.Login(username)
	if err != nil {
		return nil, errors.New("User tidak ditemukan")
	}

	if user == nil {
		return nil, errors.New("User tidak ditemukan atau tidak valid")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("password salah")
	}

	resp := &RespSuccessLogin{
		Username: user.Username,
		Name:     user.Name,
		Email:    user.Email,
		Role:     user.Role,
	}
	return resp, nil
}
