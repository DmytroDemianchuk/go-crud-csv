#  Application for request an external CSV file with a list of products
## I used the following concepts during development:
- Building a Web Application with Go, following the REST API design.
- The Clean Architecture approach in building the structure of an application. Dependency injection technique.
- Working with the <a href="https://github.com/gin-gonic/gin">gin-gonic/gin</a> framework.
- Working with Postgres database. Run from Docker. Generation of migration files.
- Writing SQL queries.

## Prerequisites
- go 1.20
- docker & docker-compose
- <a href="https://github.com/swaggo/gin-swagger">swag</a> (optional, used to re-generate swagger documentation)

## Run Project

Create .env file in root directory and add following values:

```
ENV=local

DB_NAME=store
DB_PASSWORD=test
DB_USER=root
DB_HOST=pg

HTTP_HOST=localhost
HTTP_PORT=8080
```

Definition migrating to database

```
make migrate-up
```

Use `make run` to build&run project

```
make run
```

Swagger

<a href="http://localhost:8080/swagger/index.html">http://localhost:8080/swagger/index.html</a>
