package utils

import (
	"crypto/aes"
	"crypto/cipher"
	randCry "crypto/rand"
	"demo/conf"
	"demo/database"
	"demo/models"
	"encoding/base64"
	"fmt"
	"io"
	"math/rand"
	"strings"
	"unicode"
)

var (
	FBClientID        string
	FbClientSecret    string
	GGClientID        string
	GGClientSecret    string
	AppleTeamID       string
	AppleClientID     string
	AppleCallbackURL  string
	AppleKeyID        string
	BeBaseUrl         string
	FeBaseUrl         string
	URLMappingCart    string
	NumberOfCharacter int
	TimeLife          int64
	FailedTimes       int
	LockTime          int64
	URLGetOTP         string
)

func GetConstant(conf conf.AppConfig) {
	FBClientID = conf.FBClientID
	FbClientSecret = conf.FbClientSecret
	GGClientID = conf.GGClientID
	GGClientSecret = conf.GGClientSecret
	AppleTeamID = conf.AppleTeamID
	AppleClientID = conf.AppleClientID
	AppleCallbackURL = conf.AppleCallbackURL
	AppleKeyID = conf.AppleKeyID
	BeBaseUrl = conf.BeBaseUrl
	FeBaseUrl = conf.FeBaseUrl

	// Social
	FBClientID = conf.FBClientID
	FbClientSecret = conf.FbClientSecret
	GGClientID = conf.GGClientID
	GGClientSecret = conf.GGClientSecret
	AppleTeamID = conf.AppleTeamID
	AppleClientID = conf.AppleClientID
	AppleCallbackURL = conf.AppleCallbackURL
	AppleKeyID = conf.AppleKeyID
	BeBaseUrl = conf.BeBaseUrl
	FeBaseUrl = conf.FeBaseUrl
	URLMappingCart = conf.URLMappingCart

	NumberOfCharacter = conf.NumberOfCharacter
	TimeLife = conf.TimeLife
	FailedTimes = conf.FailedTimes
	LockTime = conf.LockTime
	URLGetOTP = conf.URLGetOTP

}
func AutoMigration() (err error) {
	dbPublic, err := database.GetDatabase("default")
	if err != nil {
		return
	}
	// Create UUID Extension
	if _, err = dbPublic.DB().Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`); err != nil {
		return fmt.Errorf("error while creating DB extension 'uuid-ossp': %s", err)
	}
	// Create HStore Extension
	//if _, err = dbPublic.DB().Exec(`CREATE EXTENSION IF NOT EXISTS "hstore"`); err != nil {
	//	return fmt.Errorf("error while creating DB extension 'hstore': %s", err)
	//}
	//migrate table
	t := dbPublic.AutoMigrate(&models.User{}, &models.Otp{})
	return t.Error
}

// encrypt string to base64 crypto using AES
func Encrypt(key []byte, text string) string {
	// key := []byte(keyText)
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(randCry.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// convert to base64
	return base64.URLEncoding.EncodeToString(ciphertext)
}

// decrypt from base64 to decrypted string
func Decrypt(key []byte, cryptoText string) string {
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext)
}

// Kiem tra mat khau dung format khong
//ma ngan nhat 8 ki tu toi da 64 ki tu
//phai co chu in hoa, chu thuong va so
func VerifyPassword(password string) error {
	var uppercasePresent bool
	var lowercasePresent bool
	var numberPresent bool
	//var specialCharPresent bool
	const minPassLength = 8
	const maxPassLength = 64
	var passLen int
	var errorString string

	for _, ch := range password {
		switch {
		case unicode.IsNumber(ch):
			numberPresent = true
			passLen++
		case unicode.IsUpper(ch):
			uppercasePresent = true
			passLen++
		case unicode.IsLower(ch):
			lowercasePresent = true
			passLen++
		//case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
		//	specialCharPresent = true
		//	passLen++
		case ch == ' ':
			passLen++
		}
	}
	appendError := func(err string) {
		if len(strings.TrimSpace(errorString)) != 0 {
			errorString += ", " + err
		} else {
			errorString = err
		}
	}
	if !lowercasePresent {
		appendError("password lowercase letter missing")
	}
	if !uppercasePresent {
		appendError("password uppercase letter missing")
	}
	if !numberPresent {
		appendError("password atleast one numeric character required")
	}
	//if !specialCharPresent {
	//	appendError("special character missing")
	//}
	if !(minPassLength <= passLen && passLen <= maxPassLength) {
		appendError(fmt.Sprintf("password length must be between %d to %d characters long", minPassLength, maxPassLength))
	}

	if len(errorString) != 0 {
		return fmt.Errorf(errorString)
	}
	return nil
}
func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = LETTER_BYTES[rand.Intn(len(LETTER_BYTES))]
	}
	return string(b)
}
