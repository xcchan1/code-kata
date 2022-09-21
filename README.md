# DemystData Loan Application

## Architecture

### Backend
Application is made up of a backend service, which is able to support multiple
accounting providers and decision engines. 

### Frontend

The UI is a server rendered page, and is also served by the backend service. 
Other than the initial page load, all API calls are performed via AJAX.

## Design decisions

As different accounting providers and decision engines can have different protocol formats, 
different providers and engines can be integrated by implementing the appropriate 
`Provider` and `Engine` interfaces respectively. The implementation should convert the 
output from the provider/engine to the return format. The implemented providers
and decision engine are currently serving mocked data based on the sample provided.

## How to run

### IDE

1. Create a run configuration with the `main` function in `cmd/backend/main.go`. 
   * Optional: `PORT` environment variable can be specified, defaults to `8000`
2. Open `http://localhost:<PORT>/codekata/` in your browser

### Bash

1. Run `make build` to build the Go binary
2. Run `make run` to run the built binary. Similar to above, you may provide an optional
`PORT` environment variable
3. Open `http://localhost:<PORT>/codekata/` in your browser

### Docker

1. Run `docker-compose up --build` to build and run the service
2. Open http://localhost:8080/codekata/ in your browser

### Unit tests

1. Run `make test`
