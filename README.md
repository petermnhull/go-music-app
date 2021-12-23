# go-music-app
Simple Go API for music preferences.

## Usage
Note that this has only been tested on Ubuntu 20.04, and only intended for to be used by me as a fun project. 

### Set Up
- Create a `.env` file in the root directory and store the variables listed below.
- Download [golangci-lint](https://golangci-lint.run/usage/quick-start/) for linting.
- Download [air](https://github.com/cosmtrek/air) for live auto-reloading.

### Commands
- `make run` to run the application.
- `make build` to build the binary.
- `make test` to run the test suite.
- `make lint` to run golangci-lint.

## Variables
| Variable          | Detail                 |
|-------------------|------------------------|
| `APP_PORT`          | App server port        |
| `APP_WRITE_TIMEOUT` | App timeout for writes |
| `APP_READ_TIMEOUT`  | App timeout for reads  |
