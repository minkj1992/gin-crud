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
  - [ ] add index
  - [ ] add unique key
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
    - [ ] login
    - [ ] logout
    - [x] signup
    - [ ] find password
    - [ ] withdraw user
