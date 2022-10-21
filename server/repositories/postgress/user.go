package postgress

import (
	"errors"
	"finalproject/server/models"
	"finalproject/server/repositories"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) repositories.UserRepository {
	return &userRepo{
		db: db,
	}
}

func (a *userRepo) UserRegister(user *models.User) error {

	checkEmail := IsEmailExist(a, user.Email, 0)
	checkUsername := IsUserNameExist(a, user.Username, 0)

	if checkEmail {
		return errors.New("Email sudah terdaftar.")
	}

	if checkUsername {
		return errors.New("Username sudah terdaftar.")
	}

	err := a.db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *userRepo) UserLogin(email string) *models.User {

	var user models.User

	err := a.db.Where("Email = ?", email).Find(&user).Error
	if err != nil {
		return nil
	}

	return &user
}

func (a *userRepo) UserUpdate(user *models.User) (*models.User, error) {

	checkData := IsDataExist(a, user.Id)
	checkEmail := IsEmailExist(a, user.Email, user.Id)
	checkUsername := IsUserNameExist(a, user.Username, user.Id)

	if !checkData {
		return nil, errors.New("Data tidak ditemukan.")
	}

	if checkEmail {
		return nil, errors.New("Email sudah terdaftar.")
	}

	if checkUsername {
		return nil, errors.New("Username sudah terdaftar.")
	}

	err := a.db.Model(&user).Updates(models.User{Email: user.Email, Username: user.Username, UpdatedAt: user.UpdatedAt}).Error
	if err != nil {
		return nil, err
	}

	var dataUser models.User

	err = a.db.Where("Email = ?", user.Email).Find(&dataUser).Error
	if err != nil {
		return nil, err
	}
	return &dataUser, nil
}

func (a *userRepo) UserDelete(id int) error {
	err := a.db.Delete(&models.User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func IsEmailExist(a *userRepo, email string, id int) bool {
	var user models.User
	var err error

	if id == 0 {
		err = a.db.Where("Email = ?", email).Find(&user).Error
	} else {
		err = a.db.Where("Email = ? and Id <> ?", email, id).Find(&user).Error
	}

	if err != nil {
		return false
	}

	if user.Email == email {
		return true
	}

	return false
}

func IsUserNameExist(a *userRepo, username string, id int) bool {
	var user models.User
	var err error

	if id == 0 {
		err = a.db.Where("Username = ?", username).Find(&user).Error
	} else {
		err = a.db.Where("Username = ? and Id <> ?", username, id).Find(&user).Error
	}

	if err != nil {
		return false
	}

	if user.Username == username {
		return true
	}

	return false
}

func IsDataExist(a *userRepo, id int) bool {
	var user models.User

	err := a.db.Where("Id = ?", id).Find(&user).Error
	if err != nil {
		return false
	}

	if user.Id == id {
		return true
	}

	return false
}
