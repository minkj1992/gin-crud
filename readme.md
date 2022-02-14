# `Gin` crud server

This is a simple crud server made by `go` and `gin` framework.

## setup

```bash
go mod init github.com/minkj1992/gin-crud
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u github.com/go-sql-driver/mysql
go get -u github.com/joho/godotenv
go get -u github.com/golang-jwt/jwt
go get -u github.com/google/uuid

go mod tidy
```

> The `-u` flag instructs get to update modules providing dependencies of packages named on the command line to use newer minor or patch releases when available.

## docs

![user api](http://www.plantuml.com/plantuml/proxy?cache=no&src=https://raw.githubusercontent.com/minkj1992/gin-crud/main/docs/users.puml)

## run

```bash
go run main.go
```

## todo

- ci/cd
  - [ ] fmt script with makefile
  - [ ] ci with github action
  - [ ] cd with heroku
  - [ ] docker images
- db
  - [x] connect to mysql
  - [ ] deploy postgresql
  - [ ] 1:n relationship
  - [ ] add audio field
  - [x] add index
  - [x] add unique key
- application
  - [x] add jwt middleware
  - [x] add json validation
  - [ ] add dto
  - [ ] unittest
  - [ ] hot reload dev mode
  - [ ] add swagger
  - [x] dotenv
  - [ ] logger and sentry
  - [ ] go-admin
- features
  - [x] jwt validation
  - [ ] oauth google
  - [x] crud `TODO`
  - crud `User`
    - [x] login
    - [x] logout: jwt는 클라 또는 웹 localstorage에 저장된다 가정하면 서버에서 해줄 건 없다.
    - [x] signup
    - [x] get current user info
    - [x] get a user
    - [x] get users
    - [ ] withdraw user

## todo

### todo

#### total process

```
(http/https) requeset
-> channel -> context
-> middleware authentication -> jwt authorization
-> json validation -> dto convert
-> business model
-> dao (db connection)
-> response dto
-> response
```

### 반영할 refs

- request / response process
  - [about context process](https://www.sohamkamani.com/golang/context-cancellation-and-values/)
  - [golang blog restful gin](https://go.dev/doc/tutorial/web-service-gin)
- dto
  - [golang dto](https://stackoverflow.com/a/44981367)
  - [dto example](https://hellokoding.com/crud-restful-apis-with-go-modules-wire-gin-gorm-and-mysql/)
- dao
  - [gorm dao](https://umi0410.github.io/blog/golang/how-to-backend-in-go-db/)
