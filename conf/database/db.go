package database

import (
	"fmt"
	"sync"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DBInfo struct {
	Host       string
	Port       string
	Name       string
	User       string
	Pass       string
	SearchPath string
}

var (
	dataBaseCache = &_dbCache{cache: make(map[string]*alias)}
)

type DB struct {
	*sync.RWMutex
	DB *gorm.DB
}

type alias struct {
	Name         string
	MaxIdleConns int
	MaxOpenConns int
	DB           *DB
}

type _dbCache struct {
	mux   sync.RWMutex
	cache map[string]*alias
}

func (ac *_dbCache) add(name string, al *alias) (added bool) {
	ac.mux.Lock()
	defer ac.mux.Unlock()
	if _, ok := ac.cache[name]; !ok {
		ac.cache[name] = al
		added = true
	}
	return
}

// get database alias if cached.
func (ac *_dbCache) get(name string) (al *alias, ok bool) {
	ac.mux.RLock()
	defer ac.mux.RUnlock()
	al, ok = ac.cache[name]
	return
}

// get default alias.
func (ac *_dbCache) getDefault() (al *alias) {
	al, _ = ac.get("default")
	return
}

func GetDB(aliasNames ...string) (*gorm.DB, error) {
	var name string
	if len(aliasNames) > 0 {
		name = aliasNames[0]
	} else {
		name = "default"
	}
	al, ok := dataBaseCache.get(name)
	if ok {
		if beego.BConfig.RunMode == "dev" {
			return al.DB.DB.Debug(), nil
		}
		return al.DB.DB, nil
	}
	return &gorm.DB{}, fmt.Errorf("DataBase of alias name `%s` not found", name)
}

func RegisterDataBase(aliasName, driverName, dataSource string, params ...int) error {
	var (
		err error
		db  *gorm.DB
		al  *alias
	)

	db, err = gorm.Open(driverName, dataSource)
	if err != nil {
		err = fmt.Errorf("register database `%s`, %s", aliasName, err.Error())
		goto end
	}

	al, err = addAliasWthDB(aliasName, driverName, db)
	if err != nil {
		goto end
	}

	for i, v := range params {
		switch i {
		case 0:
			SetMaxIdleConns(al.Name, v)
		case 1:
			SetMaxOpenConns(al.Name, v)
		}
	}

end:
	if err != nil {
		if db != nil {
			//_ = db.Close()
		}
	}

	return err
}

func addAliasWthDB(aliasName, driverName string, db *gorm.DB) (*alias, error) {
	al := new(alias)
	al.Name = aliasName
	al.DB = &DB{
		RWMutex: new(sync.RWMutex),
		DB:      db,
	}

	err := db.DB().Ping()
	if err != nil {
		return nil, fmt.Errorf("register database Ping `%s`, %s", aliasName, err.Error())
	}

	if !dataBaseCache.add(aliasName, al) {
		return nil, fmt.Errorf("DataBase alias name `%s` already registered, cannot reuse", aliasName)
	}

	return al, nil
}

// get table alias.
func getDbAlias(name string) *alias {
	if al, ok := dataBaseCache.get(name); ok {
		return al
	}
	panic(fmt.Errorf("unknown DataBase alias name %s", name))
}

// SetMaxIdleConns ChangeNumber the max idle conns for *sql.DB, use specify database alias name
func SetMaxIdleConns(aliasName string, maxIdleConns int) {
	al := getDbAlias(aliasName)
	al.MaxIdleConns = maxIdleConns
	al.DB.DB.DB().SetMaxIdleConns(maxIdleConns)
}

// SetMaxOpenConns ChangeNumber the max open conns for *sql.DB, use specify database alias name
func SetMaxOpenConns(aliasName string, maxOpenConns int) {
	al := getDbAlias(aliasName)
	al.MaxOpenConns = maxOpenConns
	al.DB.DB.DB().SetMaxOpenConns(maxOpenConns)
}

func CreateDBConnectionString(info DBInfo) (dbConnString string) {
	host := info.Host
	port := info.Port
	database := info.Name
	user := info.User
	pass := info.Pass
	search_path := info.SearchPath
	dbConnString = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s search_path=%s", host, port, user, database, pass, search_path)
	return
}
