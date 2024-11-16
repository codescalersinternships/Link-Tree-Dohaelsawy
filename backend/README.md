# Link Tree ~ Backend Section

## Features:
- JWT Authentication system using cookies.
- CRUD operations for Link manegment and Profile manegment.
- Upload images to external storage - AWS S3.
- Seperate testing database from production database.
- Faster search using Redis cache db

## Project Structure:
```golang
  .
  ├── controllers // logic and handlers
  ├── database // db queries and schema 
  ├── Dockerfile.golang 
  ├── docs // swagger docs
  ├── go.mod
  ├── go.sum
  ├── main.go
  ├── middleware // control middleware
  ├── models // models and entity definition
  ├── README.md
  ├── routers // http reouters and groups 
  ├── tests // unit tests
  └── utils // helper functions 
```
## Swagger Docs:
- if you are running on local environment visit:
[http://localhost:8010/swagger/index.html#/](http://localhost:8010/swagger/index.html#/)
- if you are running on production environment visit:
[http://185.206.122.17:31010/swagger/index.html#/](http://185.206.122.17:31010/swagger/index.html#/)

## Running tests
make sure your in backend directory
```golang
cd tests
go test
```

> [!WARNING]
> before running tests, you **must explicitly** export the following environment variables
```
export AWS_REGION=****
export AWS_ACCESS_KEY_ID=****
export AWS_SECRET_ACCESS_KEY=****
export DB_CACHE_ADDR=*****
export DB_CACHE_PASSWORD=*****
``` 
