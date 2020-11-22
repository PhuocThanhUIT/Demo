package database

import (
	"strconv"
	"sync"

	"demo/conf"

	"github.com/caarlos0/env"
	"github.com/jinzhu/gorm"
)

var (
	Mutex *sync.RWMutex
)

func GetDBInfo() (dbInfo DBInfo, err error) {
	config := conf.AppConfig{}
	_ = env.Parse(&config)
	dbInfo = DBInfo{
		Host:       config.DBHost,
		Port:       config.DBPort,
		Name:       config.DBName,
		User:       config.DBUser,
		Pass:       config.DBPass,
		SearchPath: config.DBSchema,
	}
	return
}
func GetDatabase(aliasName string) (db *gorm.DB, err error) {
	var (
		customerSchema string
	)
	err = nil
	_, errConv := strconv.Atoi(aliasName)
	if errConv == nil {
		customerSchema = aliasName
	} else {
		customerSchema = "default"
	}
	db, err = GetDB(customerSchema)
	if err != nil {
		var dbInfo DBInfo
		dbInfo, err = GetDBInfo()
		if err == nil {
			err = RegisterDataBase(customerSchema, "postgres", CreateDBConnectionString(dbInfo))
			if err == nil {
				db, err = GetDB(customerSchema)
			}
		}
	}
	return
}
