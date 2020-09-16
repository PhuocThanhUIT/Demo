package utils

import (
	"demo/conf"
	"demo/database"
	"demo/models"
	"fmt"
	"math/rand"
)

var (
	FBClientID       string
	FbClientSecret   string
	GGClientID       string
	GGClientSecret   string
	AppleTeamID      string
	AppleClientID    string
	AppleCallbackURL string
	AppleKeyID       string
	BeBaseUrl        string
	FeBaseUrl        string
	URLMappingCart   string
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
	t := dbPublic.AutoMigrate(&models.User{})
	return t.Error
}
func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = LETTER_BYTES[rand.Intn(len(LETTER_BYTES))]
	}
	return string(b)
}
