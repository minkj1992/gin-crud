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

## run

```bash
go run main.go
```

# todo

- infrastructure

  - [x] add mysql
  - [x] add repository
  - [x] dotenv
  - [x] jwt auth
  - [ ] add swagger
  - [ ] ci with github action
  - [ ] cd with heroku
  - [ ] fmt script
  - [ ] hot reload dev mode

- code
  - [x] crud `entity`
  - [x] create entity relationship
    - [x] user : todo
  - [ ] crud authorization
  - [ ] add audio field to save on postgresql
  - [x] add json validation
  - [x] add middleware
  - [ ] add dto
  - [ ] unittest
