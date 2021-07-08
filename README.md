[TOC]

### API server with Golang

##### need 
```bash
    # postgresql
    # git
```

##### installed library

```bash
    # gorm - ORM library
    # sub dependency
    #      - gorm.io/driver/postgres => for database connecting
    #      - gorm.io/datatypes => for json
    # gin  - web Framework // Fiber?
    # godotenv - env file
    # dgrijalva/jwt-go - jwt generator
    # redis/v7 - Memory  // background start cli =>  redis-server --daemonize yes
    # uuid     - uuid generator
```

##### DB setting 
```bash
    # uuid extension 설치
    # CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```

##### .env 

```bash
    mkdir .env # in root directory
    
    # write key & value
    JWT_ACCESS_SECRET=YOURKEY
    JWT_REFRESH_SECRET=YOURKEY

```


##### Run App
```bash
    # dev
    go run main.go
    # prod
    go run main.go -mode="prod"
```