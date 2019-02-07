# User Service

## Getting Started

These instructions will allow you to run user service

## Requirements

Go (golang) : 1.11.2

mongodb as docker container

zipkin as docker container (optional)

## Instructions

1. Clone this repository 

2. You will notice the following directory structure

``` 
├── db.go
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
├── README.md
├── service.go
├── user-db
│   ├── Dockerfile
│   ├── seed.js
│   └── users.json
└── users.go

```

3. Set GOPATH appropriately as per the documentation - https://github.com/golang/go/wiki/SettingGOPATH
   Also, run ``` export GO111MODULE=on ```

4. Build the go application from the root of the folder

``` go build -o bin/user ```

5. Run a mongodb docker container

```sudo docker run -d -p 27017:27017 --name mgo -e MONGO_INITDB_ROOT_USERNAME=mongoadmin -e MONGO_INITDB_ROOT_PASSWORD=secret -e MONGO_INITDB_DATABASE=acmefit gcr.io/vmwarecloudadvocacy/acmeshop-user-db```

6. Export USER_HOST/USER_PORT (port and ip) as ENV variable. You may choose any used port as per your environment setup.
    
    ``` 
    export USERS_HOST=0.0.0.0
    export USERS_PORT=8081
    ```

7. Also, export ENV variables related to the database

    ```
    export USERS_DB_USERNAME=mongoadmin
    export USERS_DB_PASSWORD=secret
    export USERS_DB_HOST=0.0.0.0
    ```

8. Run the user service

```./bin/user```


### Additional Info
