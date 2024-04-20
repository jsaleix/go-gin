# go gin

Project to practice and familiarise with Golang and the Gin Web Framework.
Also implementing Server-sent-events.

## Requirements

-   `Go`
-   `Docker`

## Starting the project

First of all, start the database by running:

```
docker compose up -d
```

Then, run the following commands to start the Go app:

```
cd app/
cp .env.example .env
go get .
go run .
```

## Seeding the database

In the app directory, run one of these commands to populate the database:

```
    go run seeds/*
```

Or

```
    make seeds
```

This will create two user accounts:

| email           | Password   | Role |
| --------------- | ---------- | ----- |
| user@user.com   | User123+=  | USER  |
| admin@admin.com | Admin123+= | ADMIN |

## Routes documentation

You can use Swagger to display and try all available routes, along with the corresponding parameters and response format. The url should be http://localhost:3000/swagger/index.html unless you've changed the default port or you set the GIN_MODE to "release".

The library used to do this is [gin-swagger](https://github.com/swaggo/gin-swagger).

Note: to update swagger you'll have to run in /app `swag init`.