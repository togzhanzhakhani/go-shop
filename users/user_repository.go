package users

import (
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func(ur *UserRepository) saveUser(user *User) error {
	return ur.DB.Create(&user).Error
}

func (ur *UserRepository) getUserByEmail(email string) (*User, error) {
    var user User
    if err := ur.DB.Where("email = ?", email).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (ur *UserRepository) getAllUsers() ([]User, error) {
	var users []User
	if err := ur.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *UserRepository) GetUserByID(id string) (*User, error) {
	var user User
	if err := ur.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) updateUser(id string, updatedUser *User) error {
	if err := ur.DB.Model(&User{}).Where("id = ?", id).Updates(updatedUser).Error; err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) deleteUser(id string) error {
	if err := ur.DB.Where("id = ?", id).Delete(&User{}).Error; err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) searchUsersByName(name string) ([]User, error) {
	var users []User
	if err := ur.DB.Where("name LIKE ?", "%"+name+"%").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *UserRepository) searchUsersByEmail(email string) ([]User, error) {
	var users []User
	if err := ur.DB.Where("email = ?", email).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}