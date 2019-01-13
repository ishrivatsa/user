# User Service

## Getting Started

These instructions will allow you to run user service

## Requirements

Go (golang) : 1.11+

mongodb 

zipkin

## Instructions

1. Clone this repository 


2. You will notice the following directory structure

``` 
├── go.mod
├── go.sum
├── main.go
├── README.md
├── users
│   ├── db.go
│   ├── service.go
│   └── users.go
└── users.json
```

3. Set GOPATH appropriately as per the documentation - https://github.com/golang/go/wiki/SettingGOPATH

4. Build the go application from the root of the folder

``` go build -o bin/user ```




### Additional Info