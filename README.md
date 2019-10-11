# User Service

Version - 2.0 

## Getting Started

These instructions will allow you to run user service

## Requirements

Go (golang) : 1.11.2

mongodb as docker container

redis as docker container

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

    Expected JSON Response 
    
  	{
	    "access_token":    "eyJhbGciOiJIUzI1NiIsImtpZCI6InNpZ25pbl8xIiwidHlwIjoiSldUIn0.eyJVc2VybmFtZSI6ImVyaWMiLCJleHAiOjE1NzA3NjI5NzksInN1YiI6IjVkOTNlMTFjNmY4Zjk4YzlmYjI0ZGU0NiJ9.n70EAaiY6rbH1QzpoUJhx3hER4odW8FuN2wYG1sgH7g",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsImtpZCI6InNpZ25pbl8yIiwidHlwIjoiSldUIn0.eyJleHAiOjE1NzA3NjM1NzksInN1YiI6IjVkOTNlMTFjNmY4Zjk4YzlmYjI0ZGU0NiJ9.zwGB1340IVMLjMf_UnFC_rEeNdD131OGPcg_S0ea8DE",
    "status": 200
 	 }

   The access_token is used to make requests to other services to get data. The refresh_token is used to request new access_token. 

   If both refresh_token and access_token expire, then the user needs to log back in again. 


> **Request new access_token by using the refresh_token**
  
   **'/refresh-token methods=['POST']'**

   Expected JSON body with Request

  	 {
       "refresh_token" : "eyJhbGciOiJIUzI1NiIsImtpZCI6InNpZ25pbl8yIiwidHlwIjoiSldUIn0.eyJleHAiOjE1NzA3NjM1NzksInN1YiI6IjVkOTNlMTFjNmY4Zjk4YzlmYjI0ZGU0NiJ9.zwGB1340IVMLjMf_UnFC_rEeNdD131OGPcg_S0ea8DE"
 	  }

   Expected JSON Response

  	 {
    "access_token": "eyJhbGciOiJIUzI1NiIsImtpZCI6InNpZ25pbl8xIiwidHlwIjoiSldUIn0.eyJVc2VybmFtZSI6ImVyaWMiLCJleHAiOjE1NzA3NjMyMjksInN1YiI6IjVkOTNlMTFjNmY4Zjk4YzlmYjI0ZGU0NiJ9.wrWsDNor28aWv6huKUHAuVyROGAXqjO5luPfa5K5NQI",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsImtpZCI6InNpZ25pbl8yIiwidHlwIjoiSldUIn0.eyJleHAiOjE1NzA3NjM1NzksInN1YiI6IjVkOTNlMTFjNmY4Zjk4YzlmYjI0ZGU0NiJ9.zwGB1340IVMLjMf_UnFC_rEeNdD131OGPcg_S0ea8DE",
    "status": 200
 	 }

  Note that the access_token request here is new. 



> **Verify access_token**

   **'/verify-token' methods=['POST']**

   Expected JSON body with Request

  	 {
	"access_token": "eyJhbGciOiJIUzI1NiIsImtpZCI6InNpZ25pbl8xIiwidHlwIjoiSldUIn0.eyJVc2VybmFtZSI6ImVyaWMiLCJleHAiOjE1NzA3NjMyMjksInN1YiI6IjVkOTNlMTFjNmY4Zjk4YzlmYjI0ZGU0NiJ9.wrWsDNor28aWv6huKUHAuVyROGAXqjO5luPfa5K5NQI"
  	 }

  Expected JSON Response 

   If the the JWT is valid and user is authorized

  	{
 	   "message": "Token Valid. User Authorized",
  	   "status": 200
  	}

  If the JWT is not valid (either expired or invalid signature) then the user is NOT authorized.

  	{
    	    "message": "Invalid Key. User Not Authorized",
   	    "status": 401
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
    
