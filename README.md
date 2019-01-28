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
├── go.mod
├── go.sum
├── main.go
├── README.md
├── service.go
├── users.go
└── users.json

```

3. Set GOPATH appropriately as per the documentation - https://github.com/golang/go/wiki/SettingGOPATH

4. Build the go application from the root of the folder

``` go build -o bin/user ```

5. Run a mongodb docker container 

   ```sudo docker run -d -p 27017:27017 --name mgo -e MONGO_INITDB_ROOT_USERNAME=mongoadmin -e MONGO_INITDB_ROOT_PASSWORD=secret mongo```


6. Execute this command to import the ```users.json``` file 

   ```sudo docker cp users.json {mongodb_container_id}:/```


7. Login into the mongodb container 

    
    ```sudo docker exec -it {mongodb_container_id} bash```

8. Import the users file into the database 
    
   ```mongoimport --db acmefit --collection users --file users.json -u mongoadmin -p secret --authenticationDatabase=admin```

9. Export USER_IP/USER_PORT (port and ip) as ENV variable. You may choose any used port as per your environment setup.
    
    ``` export USERS_IP=0.0.0.0 ```
    ``` export USERS_PORT=:8087```

10. Also, export ENV variables related to the database

    ```
    export USERS_DB_USER=mongoadmin
    export USERS_DB_SECRET=secret
    export USERD_DB_IP=0.0.0.0
    ```

10. Run the user service 
  
   ```./bin/user```



### Additional Info
