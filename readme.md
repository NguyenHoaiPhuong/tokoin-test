# tokoin-simple-test

## Introduction

A simple code challenge from tokoin.

2 Solutions are aimed to provide: one programmed by golang and another one programmed by python.

## Usage

### Golang

#### Prerequisite

golang and dep (dependency management) must be installed.

#### How to run the app

- Access to golang directory:

```
cd golang
```

- In order to run the app without building the binary file, please run following commands:

```
dep ensure
go run *.go
```

- If you want to run the app by building binary file, please run following commands:

```
dep ensure
go build -o app
./app
```

### Python

## Source code structure explanation

### Golang

The app is splitted into serveral packages:

- **resources**: it is the place where we store the default configurations ```default.toml``` which need to be parsed at the beginning. Furthermore, user, ticket and organization data is also stored in json files in resources/data directory.

- **config**: this package is responsible for parsing configurations set by environment varibles or default config files (extension can be 'json', 'toml', 'yml', etc). A singleton GlobalConfig is provided for handling app configurations and it is unique.

- **model**: contains structures defining objects used in the app, such as User, Ticket and Organization.

- **jsonfunc**: contains useful json functions (parsing json object, reading json file, etc)

- **utils**: contains other useful functions, sucn as CheckError, PrintPtrStructObject, etc

### Python