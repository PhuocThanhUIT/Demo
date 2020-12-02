package models

import (
	"demo/database"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Otp struct {
	BaseModel
	PhoneNumber string    `gorm:"not null;" json:"phone_number" validate:"required"`
	Otp         string    `gorm:"not null;" json:"otp" validate:"required"`
	FailedTimes int       `gorm:"not null;" json:"failed_times" validate:"required"`
	TimeLife    time.Time `gorm:"not null;" json:"time_life" validate:"required"`
	Type        string    `gorm:"not null;" json:"type" validate:"required"`
}

func (Otp) TableName() string {
	return "otp"
}

func CreateOtp(otp Otp) (Otp, error, int) {
	dbPublic, err := database.GetDatabase("default")
	if err != nil {
		return Otp{}, err, http.StatusInternalServerError
	}
	db := dbPublic.Create(&otp)
	if db.Error != nil {
		return otp, db.Error, http.StatusForbidden
	} else {
		return otp, nil, http.StatusOK
	}
}

func GetOneOtp(PhoneNumber string) (res Otp, err error) {
	dbPublic, err := database.GetDatabase("default")
	if err != nil {
		return Otp{}, fmt.Errorf("Get one Otp: " + err.Error())
	}
	db := dbPublic.Where("phone_number = ? AND time_life > NOW()", PhoneNumber).First(&res)
	return res, db.Error
}

func UpdateOtp(id uint64) (res Otp, err error) {
	dbPublic, err := database.GetDatabase("default")
	if err != nil {
		return Otp{}, err
	}
	db := dbPublic.Model(&res).Where("id = ?", id).Update("type", "login")
	if db.Error != nil {
		return Otp{}, db.Error
	}
	dbUpdate := dbPublic.First(&res).Where("id = ?", id)
	if dbUpdate.Error != nil {
		return Otp{}, dbUpdate.Error
	}
	return res, nil
}

func IncreaseOtpFailedTime(id uint64, count int) (res Otp, err error, statusCode int) {
	dbPublic, err := database.GetDatabase("default")
	if err != nil {
		return Otp{}, err, http.StatusForbidden
	}
	db := dbPublic.Model(&res).Where("id = ?", id).Update("failed_times", count)
	if db.Error != nil {
		log.Println(db.Error)
	}
	return res, db.Error, http.StatusForbidden
}
