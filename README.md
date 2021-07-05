[TOC]

### API server with Golang

#### need 
```bash
    # postgresql
    # git
```

#### installed library

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

#### DB setting 
```bash
    # uuid extension 설치
    # CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```

#### Table list
- User
    - 회원 관리
- Notice
    - 문항 요청 관리
- Data
    - 데이터 관리
- Logic
    - 로직 관리
- Logic-Connect
    - 그룹 로직 관리 (ex : (a = 1) or (a = 2) )
- Question-Type
    - 문항 타입
- Question-Template 
    - 문항 템플릿
- Form-Attr
    - 문항 속성 관리
- Question-Form
    - 문항 보기 관리
- Question-Extra
    - 문항 추가 사항 관리 (정적 파일, 추가 필요 코딩)
- Question-Extra
    - 문항 추가 사항 관리 (정적 파일, 추가 필요 코딩)

#### Run App
```bash
    # dev
    go run main.go
    # prod
    go run main.go -mode="prod"
```