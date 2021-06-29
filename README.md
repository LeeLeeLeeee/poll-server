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
    # gin  - web Framework
```

#### Model list

```bash
    # User - 
    
```


#### Run App
```bash
    # dev
    go run main.go
    # prod
    go run main.go -mode="prod"
```