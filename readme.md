# `Gin` crud server

This is a simple crud server made by `go` and `gin` framework.

## setup

```bash
go mod init github.com/minkj1992/gin-crud
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
```

> The `-u` flag instructs get to update modules providing dependencies of packages named on the command line to use newer minor or patch releases when available.

## run

```bash
go run main.go
```

# todo

- infrastructure
  - [ ] add postgresql
  - [ ] add swagger
  - [ ] add repository
  - [ ] ci with github action
  - [ ] cd with heroku
  - [ ] fmt script
- core
  - [ ] change domain to Diary / Post
  - [ ] add audio field to save on postgresql
  - [ ] add json validation
  - [ ] add middleware
  - [ ] add dto
  - [ ] add 1:n relationship with `Diary` and `Post`
  - [ ] unittest
