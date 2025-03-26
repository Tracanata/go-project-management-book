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

func (repo *userRepository) GetDataUser(username string) (*User, error) {
	var user User
	sql := "SELECT name, email, username, role FROM user WHERE username = ?"
	result := repo.db.Raw(sql, username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// Session
// func (repo *userRepository) GetActiveSession(username string) (*Session, error) {
// 	repo.DestroySession()

// 	var session Session
// 	err := repo.db.Where("username = ? AND expires_at > ?", username, time.Now()).First(&session).Error
// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return nil, nil
// 		}
// 		return nil, err
// 	}

// 	return &session, nil
// }

// func (repo *userRepository) SetSession(username, token string, duration time.Duration) error {
// 	expirationTime := time.Now().Add(duration)
// 	session := Session{
// 		Username:   username,
// 		Token:      token,
// 		Expires_at: expirationTime,
// 	}
// 	return repo.db.Create(&session).Error
// }

// func (repo *userRepository) DestroySession() {
// 	repo.db.Where("expires_at <= NOW()").Delete(Session{})
// }
