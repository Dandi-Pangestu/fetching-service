# Project Title

This project about fetching data from the resources. 

## Getting Started

In this project there are two services, namely Auth service and Data Fetching service. Auth service is made using NodeJS and Fetching Data is created using Golang. Each will be explained above how to install it.

### Prerequisites

* [NodeJS](https://nodejs.org/en/) - Framework for Auth Service
* [NPM](https://www.npmjs.com/) - Package manager for NodeJS
* [MongoDB](https://www.mongodb.com/) - Database used

### Installing (Auth Service)

This step for installing Auth Service

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