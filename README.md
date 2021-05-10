# `shepherd-backend`

## Running the Program

Assuming you have a go environment, you can run this by doing

```
go run cmd/main.go
```

The application by default will seed the data store with 2 forms.

By default it will run a server that listens on http://localhost:8080

#### Changing Ports

```
go run cmd/main.go --port=3000
```

#### Enabling Auth

For the sake of ease of testing this server, the cookie based auth is disabled until you opt in by providing the `--auth=true` flag:

```
go run cmd/main.go --auth=true
```

## Directories

- `api`: contains all the HTTP/REST/Routing implementations
- `auth`: contains a Middleware which looks for the X-Auth cookie to have a value of `"shepherd"`
- `cmd`: contains the executable binary
- `data`: contains data persistence drivers that implement the interface described in `data/interface.go`
  - `data/mem`: contains an in memory implementation of the "database" interface
  - `data/redis`: TODO: contains a redis implementation of the interface
- `models`: describes the data entities used by this application, also describes their JSON serialization/de-serialization formats.
