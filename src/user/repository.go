package user

import "gorm.io/gorm"

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) RegisterUser(user User) error {
	result := u.db.Create(&user)
	return result.Error
}
