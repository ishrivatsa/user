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

5. Run a mongodb docker container 

   ```sudo docker run -d -p 27017:27017 --name mgo -e MONGO_INITDB_ROOT_USERNAME=mongoadmin -e MONGO_INITDB_ROOT_PASSWORD=secret mongo```


6. Execute this command to import the ```users.json``` file 

   ```sudo docker cp users.json {container_id}:/```
    
   ```mongoimport --db acmefit --collection users --file users.json -u mongoadmin -p secret --authenticationDatabase=admin```

7. Run the user service 
  
   ```./bin/user```

### Additional Info
