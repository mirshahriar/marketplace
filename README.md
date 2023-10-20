# Product Management

![Workflow](https://github.com/mirshahriar/marketplace/actions/workflows/ci.yml/badge.svg)

This application is used to manage products for a marketplace.

## Getting Started

### Prerequisites

- [Golang](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/install/)


### Local Run

To test the application locally, you can run the following commands:

#### Run MySQL

```bash
# Following command will start a MySQL container with root password as `test`
$ make docker

# If failed to create marketplace database, run the following command
$ docker exec marketplace-mysql mysql -uroot -ptest -e "CREATE DATABASE IF NOT EXISTS marketplace";
```

#### Migrate Database

```bash
# Following command will create tables in marketplace database
$ make migrate
```

For testing, run the following command to create fake user with Bearer Token.

```bash
# Following command will create a fake user with Bearer Token
$ make fake-user

2023/10/20 04:04:30 Fake user created successfully
2023/10/20 04:04:30 ID:  1
2023/10/20 04:04:30 Token:  kJVCTYxVTEAsPksG

# The generated token is the Bearer Token for the user with ID 1
```

#### Run Application

```bash
# Following command will start the application
$ make run

2023/10/20 04:05:30 Server started at port 8080
```

### Test locally

Unittests are written in the HTTP handler layer with sqlite.

To run the tests, run the following command:
```bash
# Following command will run unit tests
$ make test
```

### Deploy to Kubernetes

#### Build Docker Image

```bash
# Following command will build the docker image
$ ./build-and-push.sh

Successfully built adf253078e53
Successfully tagged mirshahriar/marketplace:latest
```

Migrate the database by running the following command:
```bash
# Following command will deploy migration job
$ ./scripts/migration.sh
```

Deploy the application to Kubernetes cluster by running the following command:
```bash
# Following command will deploy the application
$ ./scripts/deploy.sh
```

> Note: To run this application in Kubernetes, you need to have setup consul for the configuration.


## API Documentation

#### Swagger

Swagger documentation is available at `/swagger/index.html` endpoint.

```bash
# Following command will start the application
$ make run

2023/10/20 04:05:30 Server started at port 8080

# Open the following URL in browser
http://localhost:8080/swagger/index.html
```

#### Postman

Postman collection is available at `postman` directory.

## User Guide

Please check [documentation](documentation/README.md) for more details.
 

## For Developers

This application is developed using [Hexagonal Architecture](https://en.wikipedia.org/wiki/Hexagonal_architecture_(software)).

- [echo/v4](github.com/labstack/echo/v4) is used as the HTTP framework.
- [gorm.io/gorm](gorm.io/gorm) is used as the database ORM.
- [consul](https://www.consul.io/) is used for configuration management.
- [swagger](github.com/swaggo/echo-swagger) is used for API documentation.

#### To check Code Style
```bash
$ make check
```

#### To generate Swagger
```bash
$ make swag
```

#### To run unittests
```bash
$ make test
```

#### For complete checking
```bash
$ make prepare
```
    

## Check List

- [X] Implement Basic operations on products
- [X] Added validation with specific error message
- [X] Proper error handling with error codes & messages
- [X] Added unittests in the HTTP handler layer
- [X] Added user authentication with Bearer Token
- [X] Added logging middleware for HTTP requests & responses body
- [X] Added swagger documentation
- [X] Docker build support
- [X] Kubernetes deployment support
- [X] Added GoDoc for Types & Interface
- [X] Added user guide
