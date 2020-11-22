package models

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"demo/database"

	"golang.org/x/crypto/bcrypt"
)

var (
	UserList map[string]*User
)

type LoginInput struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
type User struct {
	BaseModel
	PhoneNumber string  `gorm:"size:50;not null;unique" json:"phone_number"`
	Password    *string `gorm:"size:60;not null" json:"password"`
}

func GetUserFromPhone(phone_number string) (u User, err error) {
	dbPublic, err := database.GetDatabase("default")
	if err != nil {
		return u, err
	}
	db := dbPublic.Where("phone_number = ?", phone_number).First(&u)
	return u, db.Error
}
func CheckLogin(input LoginInput) (u User, err error, statusCode int) {
	statusCode = http.StatusOK
	if input.Password == "" || input.PhoneNumber == "" {
		return u, errors.New("Empty input"), http.StatusForbidden
	}
	// GET USER FROM EMAIL
	u, err = GetUserFromPhone(input.PhoneNumber)
	if err != nil {
		return User{}, err, http.StatusForbidden
	}
	log.Println(u)
	// Comparing the password with the hash
	if *u.Password == input.Password {
		err = nil
	} else {
		err = bcrypt.CompareHashAndPassword([]byte(*u.Password), []byte(input.Password))
	}
	if err == nil {
		if err != nil {
			return u, fmt.Errorf("Error update last login: " + err.Error()), statusCode
		}
		u.Password = nil
		return u, nil, statusCode
	} else {
		return User{}, err, http.StatusInternalServerError
	}
}
func GetUserFromID(id uint64) (u User, err error) {
	log.Println(1)
	dbPublic, err := database.GetDatabase("default")
	if err != nil {
		return u, err
	}
	db := dbPublic.Where("id = ?", id).First(&u)
	return u, db.Error
}
func GetUserSocialFromID(IDSocial uint64) (user User, err error) {
	dbPublic, err := database.GetDatabase("default")
	if err != nil {
		return user, err
	}
	db := dbPublic.Where("id_social = ?", IDSocial).First(&user)
	return user, db.Error
}
func CheckExistsUserFromPhone(phone string) (res User, err error, status int) {
	dbPublic, err := database.GetDatabase("default")
	if err != nil {
		return res, err, http.StatusInternalServerError
	}
	tx := dbPublic.Model(&User{}).Where("phone_number = ?", phone).First(&res)
	if tx.RowsAffected == 0 {
		log.Println(res, err)
		return res, nil, http.StatusOK
	} else {
		return res, fmt.Errorf("So dien thoai da dang ky"), http.StatusForbidden
	}
}
func GetUserSocialFromEmail(email string) (cus User, err error) {
	dbPublic, err := database.GetDatabase("default")
	if err != nil {
		return cus, err
	}
	db := dbPublic.Where("email = ?", email).First(&cus)
	return cus, db.Error
}

func CreateUserSocial(user User) (User, error, int) {
	dbPublic, err := database.GetDatabase("default")
	fmt.Println("user ", user)
	if err != nil {
		return User{}, err, http.StatusInternalServerError
	}
	db := dbPublic.Create(&user)
	return user, db.Error, http.StatusOK
}
