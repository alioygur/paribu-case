# Classroom

Classroom is a students and teachers management system.

## Installation

If you have Go installed and configured, you can install Classroom with the following command:

the minimum version of Go required is 1.11. If you don't have Go installed, you can download the latest version from [here](https://golang.org/dl/).

before Go 1.17

```bash
go get github.com/alioygur/paribu-case/cmd/classroom
```

after Go 1.17 (with Go modules)

```bash
go install github.com/alioygur/paribu-case/cmd/classroom@latest
```

There are also pre-built binaries available for download on the `/bin` directory of this repository.

## Development

### Running the app

To run the app locally, you need to have [Go](https://golang.org/) installed.

Then, run the following commands:

```bash
$ make run
```

### Running the tests

To run the tests, run the following command:

```bash
$ make test
```

### Building the app

To build the app, run the following command:

```bash
$ make build
```
