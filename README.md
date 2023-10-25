# E-WALLET API V1

## Technology

- Golang `go 1.19`
- PostgresSQL
- Redis

## How to Start Develop
Before developing this project, you need to do some setup
1. Setup .env.yml add value from your local source
2. Run database migrations
    Follow the database migration section below

## Database Migration
We use migrations tools [goose](https://github.com/pressly/goose)

```$command
// Create database migrations file
$ migrate create -ext sql -dir migrations create_something_table

// Up the migrations
$ migrate -database "postgres://postgres:postgres@localhost:5432/db_name sslmode=disable" -path config/database/postgres up

// Down the migrations by one
$ migrate -database "postgres://postgres:postgres@localhost:5432/db_name sslmode=disable" -path config/database/postgres down
```

## Install All Package
```$command
$ make install
```

## Run Test

To run the test, just type like below

```$command
$ make test
```

## Run HTTP API
```$command
$ make run
```

But when there is any new feature, please register the path into `Makefile`
