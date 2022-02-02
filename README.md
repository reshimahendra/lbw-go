# LBW-GO, GIN Restful-api Project starter

[Golang][1] restful-api project starter-kit with [Gin][11] and [Postgresql][2] (using [pgx][12] library). This starter project build using `DDD` model architect.

```bash
# NOTE: PROJECT is NOT READY yet.
```

## Table of Content
1. [Quick Review](#1-quick-review)
2. [Directory Structure](#2-directory-structure)
3. [Getting started](#3-getting-started)
4. [TODO](#4-todo)

### 1. Quick Review
The package used in this project:
- [x] [Gin](11), most stared [Go](1) web framework on [Github](3)
- [x] [pgx](12), the best [PostgreSQL](2) driver for [Go](1)
- [x] [jwt](16), Jwt Authentication and Authorization
- [x] [logrus](13), logger for access log, database log, and server log
- [x] [viper](14), complete configuration solution
- [x] [Air](15), Hot reload support for fast development

### 2. Directory Structure
```bash
|-- root/
|-- |-- cmd/
|-- |-- |-- app/
|-- |-- |-- |-- server/
|-- |-- |-- main.go
|-- |-- config/
|-- |-- |-- .config.yaml
|-- |-- dist/
|-- |-- internal/
|-- |-- |-- app/
|-- |-- |-- |-- account/
|-- |-- |-- |-- |-- datastore/
|-- |-- |-- |-- |-- handler/
|-- |-- |-- |-- |-- service/
|-- |-- |-- |-- |-- README.md
|-- |-- |-- |-- |-- router.go
|-- |-- |-- |-- mail/
|-- |-- |-- config/
|-- |-- |-- database/
|-- |-- |-- |-- sql/
|-- |-- |-- domain/
|-- |-- |-- pkg/
|-- |-- |-- |-- auth/
|-- |-- |-- |-- errors/
|-- |-- |-- |-- helper/
|-- |-- |-- |-- logger/
|-- |-- log/
|-- |-- vendor/
|-- |-- go.mod
|-- |-- go.sum
|-- |-- LICENSE
|-- |-- Makefile
|-- |-- README.md
```
### 3. Getting Started

#### copy config file

Before we run the app, make sure the config has been set.
from the `root` directory, run
```bash
cp config/example.config.yaml .config.yaml
```

#### run app

to test application (with hot reload), run:
```bash
air
```

OR

run without hot reload:
```bash
make run
```

#### build app

To build the app, run:
```bash
make build
```

#### run test

To run the test, run:
```bash
make test
```

Show test coverage result on browser:
```bash
# make sure file 'proof.out already exist/ generated'
make show-test
```

Generate new test coverage and show result on browser:
```bash
make test-proof
```

### 4. TODO

- [x] Jwt Authentication and Authorization
- [x] Add Test
- [ ] Add blog app


[1]:https://golang.org
[2]:https://www.postgresql.org
[3]:https://github.com
[11]:https://github.com/gin-gonic/gin
[12]:https://github.com/jackc/pgx
[13]:https://github.com/sirupsen/logrus
[14]:https://github.com/spf13/viper
[15]:https://github.com/cosmtrek/air
[16]:https://github.com/golang-jwt/jwt
