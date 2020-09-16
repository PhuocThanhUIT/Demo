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

type User struct {
	BaseModel
	Email    string  `gorm:"size:50;not null;unique" json:"email"`
	Password *string `gorm:"size:60;not null" json:"password"`
}

func GetCustomerFromEmail(email string) (u User, err error) {
	dbPublic, err := database.GetDatabase("default")
	if err != nil {
		return u, err
	}
	db := dbPublic.Where("email = ?", email).First(&u)
	return u, db.Error
}
func CheckLogin(input LoginInput) (u User, err error, statusCode int) {
	statusCode = http.StatusOK
	if input.Password == "" || input.Email == "" {
		return u, errors.New("Empty input"), http.StatusForbidden
	}
	// GET USER FROM EMAIL
	u, err = GetCustomerFromEmail(input.Email)
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
func GetCustomerFromID(id uint64) (u User, err error) {
	log.Println(1)
	dbPublic, err := database.GetDatabase("default")
	if err != nil {
		return u, err
	}
	db := dbPublic.Where("id = ?", id).First(&u)
	return u, db.Error
}
func GetCustomerSocialFromID(IDSocial uint64) (customer User, err error) {
	dbPublic, err := database.GetDatabase("default")
	if err != nil {
		return customer, err
	}
	db := dbPublic.Where("id_social = ?", IDSocial).First(&customer)
	return customer, db.Error
}

func GetCustomerSocialFromEmail(email string) (cus User, err error) {
	dbPublic, err := database.GetDatabase("default")
	if err != nil {
		return cus, err
	}
	db := dbPublic.Where("email = ?", email).First(&cus)
	return cus, db.Error
}

func CreateCustomerSocial(customer User) (User, error, int) {
	dbPublic, err := database.GetDatabase("default")
	fmt.Println("customer ", customer)
	if err != nil {
		return User{}, err, http.StatusInternalServerError
	}
	db := dbPublic.Create(&customer)
	return customer, db.Error, http.StatusOK
}
