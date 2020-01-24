# User

[![gcr.io](https://img.shields.io/badge/gcr.io-stable-green?style=flat-square)](https://console.cloud.google.com/gcr/images/vmwarecloudadvocacy/GLOBAL/acmeshop-user@sha256:9853344bb335df2900de8d5454ada13422de62c8b72852564688d522fda9290d/details?tab=info)

> A user service, because what is a shop without users to buy our awesome red pants?

The User service is part of the [ACME Fitness Shop](https://github.com/vmwarecloudadvocacy/acme_fitness_demo). The goal of this specific service is to register and authenticate users using JWT tokens.

## Prerequisites

There are different dependencies based on whether you want to run a built container, or build a new one.

### Build

* [Go (at least Go 1.12)](https://golang.org/dl/)
* [Docker](https://www.docker.com/docker-community)

### Run

* [Docker](https://www.docker.com/docker-community)
* [MongoDB](https://hub.docker.com/r/bitnami/mongodb)
* [Redis](https://hub.docker.com/r/bitnami/redis)
* [Zipkin](https://hub.docker.com/r/openzipkin/zipkin)

## Installation

### Docker

Use this command to pull the latest tagged version of the shipping service:

```bash
docker pull gcr.io/vmwarecloudadvocacy/acmeshop-user:stable
```

To build a docker container, run `docker build . -t vmwarecloudadvocacy/acmeshop-user:<tag>`.

The images are tagged with:

* `<Major>.<Minor>.<Bug>`, for example `1.1.0`
* `stable`: denotes the currently recommended image appropriate for most situations
* `latest`: denotes the most recently pushed image. It may not be appropriate for all use cases

### Source

To build the app as a stand-alone executable, run `go build` from the `cmd/users` directory.

## Usage

The **user** service, either running inside a Docker container or as a stand-alone app, relies on the below environment variables:

* **USERS_HOST**: The IP of the user app to listen on (like `0.0.0.0`)
* **USERS_PORT**: The port number for the user service to listen on (like `8083`)
* **USERS_DB_USERNAME**: The username to connect to the MongoDB database
* **USERS_DB_PASSWORD**: The password to connect to the MongoDB database
* **USERS_DB_HOST**: The host or IP on which MongoDB is active
* **USERS_DB_PORT**: The port on which MongoDB is active
* **REDIS_HOST**: The host or IP on which RedisDB is active
* **REDIS_PORT**: The port on which RedisDB is active
* **REDIS_PASSWORD**: The password for RedisDB. This field must be provided else the value defaults to *secret*
* **JAEGER_AGENT_HOST**: The host for Jaeger agent - Use this only if you want tracing enabled
* **JAEGER_AGENT_PORT**: The port for Jaeger agent - Use this only if you want tracing enabled

The Docker image is based on the Bitnami MiniDeb container. Use this commands to run the latest stable version of the payment service with all available parameters:

```bash
# Run the MongoDB container
docker run -d -p 27017:27017 --name mgo -e MONGO_INITDB_ROOT_USERNAME=mongoadmin -e MONGO_INITDB_ROOT_PASSWORD=secret -e MONGO_INITDB_DATABASE=acmefit gcr.io/vmwarecloudadvocacy/acmeshop-user-db:latest

# Run the Redis container
docker run -d -p 6379:6379 -e REDIS_PASSWORD=secret --name redis bitnami/redis

# Run the user service
docker run -d -e USERS_HOST=0.0.0.0 -e USERS_PORT=8081 -e USERS_DB_USERNAME=mongoadmin -e USERS_DB_PASSWORD=secret -e USERS_DB_HOST=0.0.0.0 -e REDIS_HOST=0.0.0.0 -e REDIS_PORT=6379 -e REDIS_PASSWORD=secret -e JAEGER_AGENT_HOST=localhost -e JAEGER_AGENT_PORT=6831 -p 8083:8083 gcr.io/vmwarecloudadvocacy/acmeshop-user:stable
```

## Available users

There are four pre-created users loaded into the database:

| User   | Password   |
|--------|------------|
| eric   | `vmware1!` |
| dwight | `vmware1!` |
| han    | `vmware1!` |
| phoebe | `vmware1!` |

## API

### HTTP

#### `GET /users`

Returns the list of all users

```bash
curl --request GET \
  --url http://localhost:8083/users \
   -H 'Authorization: Bearer <TOKEN>'
```

```json
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
```

#### `GET /users/:id`

Returns details about a specific user id

```bash
curl --request GET \
  --url http://localhost:8083/users/5c61ed848d891bd9e8016899 \
    -H 'Authorization: Bearer <TOKEN>'
  
```

```json
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
```

#### `POST /login/`

Authenticate and Login user

```bash
curl --request POST \
  --url http://localhost:8083/login \
  --header 'content-type: application/json' \
  --data '{ 
    "username": "username",
    "password": "password"
}'
```

The request to login needs to have a username and password

```json
{ 
    "username": "username",
    "password": "password"
}
```

When the login succeeds, an access token is returned

```json
{
    "access_token":    "eyJhbGciOiJIUzI1NiIsImtpZCI6InNpZ25pbl8xIiwidHlwIjoiSldUIn0.eyJVc2VybmFtZSI6ImVyaWMiLCJleHAiOjE1NzA3NjI5NzksInN1YiI6IjVkOTNlMTFjNmY4Zjk4YzlmYjI0ZGU0NiJ9.n70EAaiY6rbH1QzpoUJhx3hER4odW8FuN2wYG1sgH7g",
"refresh_token": "eyJhbGciOiJIUzI1NiIsImtpZCI6InNpZ25pbl8yIiwidHlwIjoiSldUIn0.eyJleHAiOjE1NzA3NjM1NzksInN1YiI6IjVkOTNlMTFjNmY4Zjk4YzlmYjI0ZGU0NiJ9.zwGB1340IVMLjMf_UnFC_rEeNdD131OGPcg_S0ea8DE",
"status": 200
    }
```

The access_token is used to make requests to other services to get data. The refresh_token is used to request new access_token. If both refresh_token and access_token expire, then the user needs to log back in again.

#### `POST /refresh-token`

Request new access_token by using the `refresh_token`

```bash
curl --request POST \
  --url http://localhost:8083/refresh-token \
  --header 'content-type: application/json' \
  --data '{
    "refresh_token" : "eyJhbGciOiJIUzI1NiIsImtpZCI6InNpZ25pbl8yIiwidHlwIjoiSldUIn0.eyJleHAiOjE1NzA3NjM1NzksInN1YiI6IjVkOTNlMTFjNmY4Zjk4YzlmYjI0ZGU0NiJ9.zwGB1340IVMLjMf_UnFC_rEeNdD131OGPcg_S0ea8DE"
}'
```

The request to the refresh-token service, needs a valid refresh_token

```json
{
    "refresh_token" : "eyJhbGciOiJIUzI1NiIsImtpZCI6InNpZ25pbl8yIiwidHlwIjoiSldUIn0.eyJleHAiOjE1NzA3NjM1NzksInN1YiI6IjVkOTNlMTFjNmY4Zjk4YzlmYjI0ZGU0NiJ9.zwGB1340IVMLjMf_UnFC_rEeNdD131OGPcg_S0ea8DE"
}
```

When the token is valid, a new access_token is returned

```json
{
    "access_token": "eyJhbGciOiJIUzI1NiIsImtpZCI6InNpZ25pbl8xIiwidHlwIjoiSldUIn0.eyJVc2VybmFtZSI6ImVyaWMiLCJleHAiOjE1NzA3NjMyMjksInN1YiI6IjVkOTNlMTFjNmY4Zjk4YzlmYjI0ZGU0NiJ9.wrWsDNor28aWv6huKUHAuVyROGAXqjO5luPfa5K5NQI",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsImtpZCI6InNpZ25pbl8yIiwidHlwIjoiSldUIn0.eyJleHAiOjE1NzA3NjM1NzksInN1YiI6IjVkOTNlMTFjNmY4Zjk4YzlmYjI0ZGU0NiJ9.zwGB1340IVMLjMf_UnFC_rEeNdD131OGPcg_S0ea8DE",
    "status": 200
}
```

#### `POST /verify-token`

Verify access_token

```bash
curl --request POST \
  --url http://localhost:8083/verify-token \
  --header 'content-type: application/json' \
  --data '{
    "access_token": "eyJhbGciOiJIUzI1NiIsImtpZCI6InNpZ25pbl8xIiwidHlwIjoiSldUIn0.eyJVc2VybmFtZSI6ImVyaWMiLCJleHAiOjE1NzA3NjMyMjksInN1YiI6IjVkOTNlMTFjNmY4Zjk4YzlmYjI0ZGU0NiJ9.wrWsDNor28aWv6huKUHAuVyROGAXqjO5luPfa5K5NQI"
}'
```

The request to verify-token needs a valid access_token

```json
{
    "access_token": "eyJhbGciOiJIUzI1NiIsImtpZCI6InNpZ25pbl8xIiwidHlwIjoiSldUIn0.eyJVc2VybmFtZSI6ImVyaWMiLCJleHAiOjE1NzA3NjMyMjksInN1YiI6IjVkOTNlMTFjNmY4Zjk4YzlmYjI0ZGU0NiJ9.wrWsDNor28aWv6huKUHAuVyROGAXqjO5luPfa5K5NQI"
}
```

If the the JWT is valid and user is authorized, an HTTP/200 message is returned

```json
{
   "message": "Token Valid. User Authorized",
   "status": 200
}
```

If the JWT is not valid (either expired or invalid signature) then the user is NOT authorized and an HTTP/401 message is returned

```json
{
    "message": "Invalid Key. User Not Authorized",
    "status": 401
}
```

#### `POST /logout`

Logout the user - Adds the access token to redis cache

```
curl -X POST \
  http://localhost:8083/logout \
  -H 'Authorization: Bearer <TOKEN>' 
```


#### `POST /register`

Register/Create new user

```bash
curl --request POST \
  --url http://localhost:8083/register \
  --header 'content-type: application/json' \
  --data '{
    "username":"peterp",
    "password":"vmware1!",
    "firstname":"amazing",
    "lastname":"spiderman",
    "email":"peterp@acmefitness.com"
}'
```

To create a new user, a valid user object needs to be provided

```json
{
    "username":"peterp",
    "password":"vmware1!",
    "firstname":"amazing",
    "lastname":"spiderman",
    "email":"peterp@acmefitness.com"
}
```

When the user is successfully created, an HTTP/201 message is returned

```json
{
    "message": "User created successfully!",
    "resourceId": "5c61ef891d41c8de20281dd2",
    "status": 201
}
```

## Contributing

[Pull requests](https://github.com/vmwarecloudadvocacy/user/pulls) are welcome. For major changes, please open [an issue](https://github.com/vmwarecloudadvocacy/user/issues) first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

See the [LICENSE](./LICENSE) file in the repository
