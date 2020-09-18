module github.com/PhuocThanhUIT/Demo

go 1.15
import (
	demo/conf 
	demo/utils
	demo/models
	demo/controllers
	demo/routes
	demo/database
	demo/helper
)
require (
	github.com/BillSJC/appleLogin v0.0.0-20190916123152-090c3039745d
	github.com/astaxie/beego v1.12.2
	github.com/caarlos0/env v3.5.0+incompatible
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-redis/redis/v7 v7.4.0
	github.com/go-redis/redis/v8 v8.0.0-beta.11
	github.com/iris-contrib/go.uuid v2.0.0+incompatible
	github.com/jinzhu/gorm v1.9.16
	github.com/lib/pq v1.8.0 // indirect
	github.com/parnurzeal/gorequest v0.2.16 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	golang.org/x/net v0.0.0-20200904194848-62affa334b73 // indirect
	golang.org/x/oauth2 v0.0.0-20200902213428-5d25da1a8d43
	moul.io/http2curl v1.0.0 // indirect
)
