package user

import (
	"errors"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (repo *userRepository) Register(user *User) error {
	result := repo.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *userRepository) Login(username string) (*User, error) {
	var user User
	sql := "SELECT name, email, password, username, role FROM user WHERE username = ?"
	result := repo.db.Raw(sql, username).Scan(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("username tidak ditemukan")
		}
		return nil, result.Error
	}

	return &user, nil
}
