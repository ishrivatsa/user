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
├── cmd
│   └── users
│       └── main.go
├── Dockerfile
├── entrypoint
│   └── docker-entrypoint.sh
├── go.mod
├── go.sum
├── internal
│   ├── auth
│   │   └── auth.go
│   ├── db
│   │   └── db.go
│   └── service
│       └── service.go
├── pkg
│   └── logger
│       └── logger.go
├── README.md
└── user-db
    ├── Dockerfile
    ├── seed.js
    └── users.json

```

3. Set GOPATH appropriately as per the documentation - https://github.com/golang/go/wiki/SettingGOPATH
   Also, run ``` export GO111MODULE=on ```

4. Build the go application from the root of the folder

   ``` go build -o bin/user ./cmd/users```

5. Run a mongodb docker container

  ```sudo docker run -d -p 27017:27017 --name mgo -e MONGO_INITDB_ROOT_USERNAME=mongoadmin -e      MONGO_INITDB_ROOT_PASSWORD=secret -e MONGO_INITDB_DATABASE=acmefit gcr.io/vmwarecloudadvocacy/acmeshop-user-db```

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


## Additional Info 
   
There are pre-created users loaded into the database. 

**Username(s): eric, dwight, han, phoebe**  

**Password: vmware1!  **


## API

> **Returns the list of all users**
   
   **'/users' methods=['GET']**

    Expected JSON Response 

    {
    "data": [
        {
            "username": "walter",
            "email": "walter@acmefitness.com",
            "firstname": "Walter",
            "lastname": "White",
            "id": "5c61ed848d891bd9e8016898"
        },
        {
            "username": "dwight",
            "email": "dwight@acmefitness.com",
            "firstname": "Dwight",
            "lastname": "Schrute",
            "id": "5c61ed848d891bd9e8016899"
        }
    ]}
    


> **Returns details about a specific user id**
   
   **'/users/:id' methods=['GET']**

    Expected JSON response

    {
        "data": {
            "username": "dwight",
            "email": "dwight@acmefitness.com",
            "firstname": "Dwight",
            "lastname": "Schrute",
            "id": "5c61ed848d891bd9e8016899"
        },
        "status": 200
    }


> **Authenticate and Login user**
   
   **'/login/' methods=['POST']**

    Expected JSON Body with the request
     
     { 
           "username": "username",
           "password": "password"
     }

    Expected JSON Response - The token here is JWT. This has to be sent as ```Authorization Bearer <TOKEN>``` token with every subsequent request.

    
    {
        "status": 200,
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImVyaWMiLCJleHAiOjE1NzAyMjkyOTl9.64S2amtZ4-cp0kbrGLqexyMJ4f0PrnSXKMc-xvNqoqs"
    }

> **Register/Create new user**

   **'/register' methods=['POST']**

    Expected JSON body with Request

    {
    	"username":"peterp",
    	"password":"vmware1!",
    	"firstname":"amazing",
    	"lastname":"spiderman",
    	"email":"peterp@acmefitness.com"
    }
    

    Expected JSON Response 

    
    {
        "message": "User created successfully!",
        "resourceId": "5c61ef891d41c8de20281dd2",
        "status": 201
    }
    
