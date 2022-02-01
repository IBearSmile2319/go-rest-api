 #  # GO-REST-API  <img src="https://i.pinimg.com/originals/e4/97/e9/e497e9cfa0c8d4c0bfd78c2c508c6f09.gif" width="50">    

## Running Postgres Locally with Docker
```sh
$ docker run --name some-postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres
```
## Available Scripts local

first we configure the variables for the connection to the postgres database.

```sh
$ export DB_USERNAME=postgres
$ export DB_PASSWORD=postgres
$ export DB_TABLE=postgres
$ export DB_PORT=5432
$ export DB_DB=postgres
```
In the project directory, you can run:

### Run Script localStorage 

```sh
$ go run cmd/server/main.go
```
Open [http://localhost:8080/api/health] to view en Postman.
____

## Containerizing our Go Apps with Docker

How we can effectively containerize our Go applications using Docker.

```sh
$ docker build -t comments-api
```
**[optional] - It gives error when executing it but just in case.**

```sh
$ docker run -it -p 8080:8080 comments-api
```
### Docker-compose for our Go Services
We can easily spin up all of the containers for our REST API. **"docker-compose.yml"**

```sh
$ docker-compose up --build
```

---
## Acceptance Tests with Resty

```sh
$ go test ./... -tags=e2e -v
```



















