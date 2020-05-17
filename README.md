# Efishery Service

This project about fetching data from the resources. 

## Getting Started

In this project there are two services, namely Auth service and Data Fetching service. Auth service is made using NodeJS and Fetching Data is created using Golang. Each will be explained above how to install it. See api documentations [here](https://github.com/Dandi-Pangestu/efishery-service/blob/master/doc-api.md).

### Prerequisites

* [NodeJS](https://nodejs.org/en/) - Framework for Auth Service
* [NPM](https://www.npmjs.com/) - Package manager for NodeJS
* [MongoDB](https://www.mongodb.com/) - Database used
* [Go](https://golang.org/) - Framework for Fetching Service
* [Redis](https://redis.io/) - Caching data

### Installing (Auth Service)

This step for installing Auth Service

Make sure MongoDB already on
```
sudo service mongod start 
```

Prepare environment
```
cd auth
cp .env.example .env
```

Install dependencies
```
npm install
```

Running server
```
npm run dev
```

The server will running at localhost:4000

### Installing (Fetching Service)

This step for installing Fetching Service

Make sure Redis already on
```
sudo service redis-server start 
```

Prepare environment
```
cd fetching
```

Install dependencies
```
go mod tidy
```

Running server
```
go build
./micro
```

The server will running at localhost:8080

## Running the tests

There are two tests, namely testing for Auth Service and Fetching Data Service. Below, each testing will be explained.

### Testing in Auth Service

Prepare environment
```
cd auth
cp .env.test.example .env.test
```

Running test
```
npm run test
```

### Testing in Fetching Service

Prepare environment
```
cd fetching/shared/utils
```

Running test
```
go test -v
```