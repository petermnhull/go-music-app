# go-music-app
Simple Go API for music preferences.

## Usage
Note that this has only been tested on Ubuntu 20.04, and is only intended for to be used by me as a fun project. 

### Set Up
- Create a `.env` file in the root directory and store the variables listed below.
- Download [golangci-lint](https://golangci-lint.run/usage/quick-start/) for linting.
- Download [air](https://github.com/cosmtrek/air) for live auto-reloading.

### Dependencies
- Postgres database.

### Commands
- `make run` to run the application.
- `make build` to build the binary.
- `make test` to run the test suite.
- `make lint` to run golangci-lint.

### Migrations
Postgres migrations are managed with [dbmate](https://github.com/amacneil/dbmate) and can be found in the `db` directory.

### Variables
| Variable            | Detail                 | Default     |
|---------------------|------------------------|-------------|
| `APP_PORT`          | App server port        | `8080`      |
| `APP_ADDRESS`       | App server address     | `127.0.0.1` |
| `APP_WRITE_TIMEOUT` | App timeout for writes | `5s`        |
| `APP_READ_TIMEOUT`  | App timeout for reads  | `5s`        |
| `POSTGRES_URL`      | Postgres URL           | None        |            
