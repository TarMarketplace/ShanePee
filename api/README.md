<div align="center">
    <h1>
        Shanepee API
    </h1>
</div>

## Prerequisites

- [Golang 1.23](https://go.dev/dl/)
- make (should come with build tools XCode/git bash)
- [Docker](https://docs.docker.com/engine/install/) (optional)
- [Wire](https://github.com/google/wire)

## Building application

build application with `make`

```sh
make
```

To update `openapi.json` run

```sh
make docs
```

To compile `.feature` file to `.go` and execute tests run

```sh
    cd tests && go test
```

## Artchitecture

`apperror/` contains `AppError` interface. This is method allow application error to convert into HTTP error response.

`cmd/` contains `main` package which is use to wire thing together.

`config/` contains application config loader.

`docs/` generated openAPI documentation from `make docs`.

`domain/` pure data model. Since this is pure business logic, there should be no implementation detail (database, HTTP) on this layer. Currently, for ease of implementation, simple struct tag can be used here.

`infrastructure/` implementation detail.

`service/` contains usecase level code. Manipulation of domains object should be perform here.

## TODO

- proper logger
- Input validation
- DB File config
