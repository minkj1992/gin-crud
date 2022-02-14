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
    - [x] logout: jwt는 httponly-cookie에 저장된다 가정하면 서버에서 해줄 건 없다.
    - [x] signup
    - [x] get current user info
    - [x] get a user
    - [x] get users
    - [ ] withdraw user

## logout 정리

서버가 jwt로 access-token을 클라에게 줄 때, xss 공격을 방지하기 위해 HttpOnly-Cookie에 넣어서 보내준다.
이렇게 되면 "로그아웃"기능을 구현할 때 기존 cookie 또는 localstorage로 jwt를 저장하는 것과 비교했을 때, 보안적으로는 강력해지지만 단점이 클라에서 httponly-cookie를 수정하지 못하게 된다.

그렇다보니 기존 방식에서는 logout 기능을 간단하게 웹프론트/클라가 cookie를 지워주면 되었지만, httponly-Cookie에서는 서버가 logout api를 제공해주어야 한다. 이 때 api는 요청을 받게되면 httponly-cookie의 token 데이터를 expire 시켜주어야 한다.

rest가 중간에 한번 필요하기 때문에, 클라이언트가 네트워크 offline 상태일 때, 로그아웃을 못한다는 조금 이상한 상황이 연출 될 수 있다. 뭐 요즘에 네트워크 안되는 일이 거의 없긴하지만, 단점인건 분명하다. 그래서 offline 로그아웃은 어떻게 해줘야 하는지, 검색해보니 [Logging Out While Offline With HttpOnly Cookies](https://medium.com/@thbrown/logging-out-with-http-only-session-ad09898876ba) 포스팅이 있어서 이후에 정리해보고자 한다.

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
- jwt
  - [jwt-gin-redis 엣지 포인트 잘 정리된 포스트](https://learn.vonage.com/blog/2020/03/13/using-jwt-for-authentication-in-a-golang-application-dr/)
  - TODO: [Logging Out While Offline With HttpOnly Cookies](https://medium.com/@thbrown/logging-out-with-http-only-session-ad09898876ba)
